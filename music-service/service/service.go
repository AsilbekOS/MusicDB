package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"music/proto"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SongService struct {
	proto.UnimplementedSongServiceServer
	PostgreSQL *sql.DB
}

func NewSongService(db *sql.DB) *SongService {
	return &SongService{PostgreSQL: db}
}

func (m *SongService) AddSong(ctx context.Context, req *proto.AddSongRequest) (*proto.AddSongResponse, error) {
	// Validate incoming request
	if req.GetGroup() == "" || req.GetSong() == "" {
		return nil, status.Error(codes.InvalidArgument, "Group name or song name cannot be empty")
	}

	// Retrieve access token
	token, err := GetAccessToken()
	if err != nil {
		log.Printf("Failed to retrieve access token: %v", err)
		return nil, status.Error(codes.Internal, "Unable to retrieve access token")
	}

	// Search for the song
	songInfo, err := SearchSong(req.Song, token)
	if err != nil {
		log.Printf("Error while searching for the song: %v", err)
		return nil, status.Error(codes.Internal, "Error occurred during song search")
	}

	if len(songInfo.Tracks.Items) == 0 {
		log.Println("No matching song found.")
		return nil, status.Error(codes.NotFound, "No matching song found")
	}

	// Extract song information
	track := songInfo.Tracks.Items[0]
	releaseDate, err := time.Parse("2006-01-02", track.Album.ReleaseDate)
	if err != nil {
		log.Printf("Failed to parse release date: %v", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid release date format")
	}

	// Insert song into the database
	query := `INSERT INTO songs (group_name, song_name, release_date, link) VALUES ($1, $2, $3, $4) RETURNING id`
	var songID string
	err = m.PostgreSQL.QueryRowContext(ctx, query, req.GetGroup(), req.GetSong(), releaseDate, track.ExternalURLs.Spotify).Scan(&songID)
	if err != nil {
		log.Printf("Database insertion error: %v", err)
		return nil, status.Errorf(codes.Internal, "Error inserting song into the database: %v", err)
	}

	log.Println("Song successfully added to the database")
	return &proto.AddSongResponse{
		SongId: songID,
		Status: "success",
	}, nil
}

func (m *SongService) GetSong(ctx context.Context, req *proto.GetSongRequest) (*proto.GetSongResponse, error) {
	query := `SELECT id, group_name, song_name, release_date, link FROM songs WHERE id = $1`
	var song proto.GetSongResponse
	err := m.PostgreSQL.QueryRowContext(ctx, query, req.GetSongId()).Scan(&song.SongId, &song.Group, &song.Song, &song.ReleaseDate, &song.Link)
	if err != nil {
		log.Printf("Song with ID %s not found: %v", req.GetSongId(), err)
		return nil, status.Error(codes.NotFound, "Song not found in the database")
	}

	return &song, nil
}

func (m *SongService) GetSongs(ctx context.Context, req *proto.GetSongsRequest) (*proto.GetSongsResponse, error) {
	query := `SELECT id, group_name, song_name, release_date, link FROM songs WHERE group_name ILIKE $1 LIMIT $2 OFFSET $3`
	rows, err := m.PostgreSQL.QueryContext(ctx, query, "%"+req.GetGroup()+"%", req.GetPageSize(), (req.GetPage()-1)*req.GetPageSize())
	if err != nil {
		log.Printf("Error retrieving songs: %v", err)
		return nil, status.Error(codes.Internal, "Failed to retrieve songs from the database")
	}
	defer rows.Close()

	var songs []*proto.Song
	for rows.Next() {
		var song proto.Song
		if err := rows.Scan(&song.SongId, &song.Group, &song.Song, &song.ReleaseDate, &song.Link); err != nil {
			log.Printf("Error scanning song data: %v", err)
			return nil, status.Error(codes.Internal, "Failed to process song data")
		}
		songs = append(songs, &song)
	}

	return &proto.GetSongsResponse{
		Songs: songs,
	}, nil
}

func (m *SongService) UpdateSong(ctx context.Context, req *proto.UpdateSongRequest) (*proto.UpdateSongResponse, error) {
	// Retrieve access token
	token, err := GetAccessToken()
	if err != nil {
		log.Printf("Failed to retrieve access token: %v", err)
		return nil, status.Error(codes.Internal, "Unable to retrieve access token")
	}

	// Search for the song
	songInfo, err := SearchSong(req.Song, token)
	if err != nil {
		log.Printf("Error while searching for the song: %v", err)
		return nil, status.Error(codes.Internal, "Error occurred during song search")
	}

	if len(songInfo.Tracks.Items) == 0 {
		log.Println("No matching song found.")
		return nil, status.Error(codes.NotFound, "No matching song found")
	}

	// Extract song information
	track := songInfo.Tracks.Items[0]
	releaseDate, err := time.Parse("2006-01-02", track.Album.ReleaseDate)
	if err != nil {
		log.Printf("Failed to parse release date: %v", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid release date format")
	}
	songUrl := track.ExternalURLs.Spotify

	// Start building dynamic SQL query
	query := "UPDATE songs SET "
	var args []interface{}
	var setClauses []string
	var argIndex int

	// Update group name if provided
	if req.GetGroup() != "" {
		setClauses = append(setClauses, "group_name = $"+strconv.Itoa(argIndex+1))
		args = append(args, req.GetGroup())
		argIndex++
	}

	// Update song name if provided
	if req.GetSong() != "" {
		setClauses = append(setClauses, "song_name = $"+strconv.Itoa(argIndex+1))
		args = append(args, req.GetSong())
		argIndex++

		// If song name is updated, also update release date and URL
		setClauses = append(setClauses, "release_date = $"+strconv.Itoa(argIndex+1))
		args = append(args, releaseDate)
		argIndex++

		setClauses = append(setClauses, "link = $"+strconv.Itoa(argIndex+1))
		args = append(args, songUrl)
		argIndex++
	}

	// If no fields are provided to update
	if len(setClauses) == 0 {
		log.Println("No fields provided to update.")
		return nil, status.Error(codes.InvalidArgument, "No fields provided for update")
	}

	// Complete the SQL query by adding the WHERE clause for ID
	query += fmt.Sprintf("%s WHERE id = $%d", strings.Join(setClauses, ", "), argIndex+1)
	args = append(args, req.GetSongId())

	// Execute the query
	_, err = m.PostgreSQL.ExecContext(ctx, query, args...)
	if err != nil {
		log.Printf("Error updating song with ID %s: %v", req.GetSongId(), err)
		return nil, status.Error(codes.Internal, "Failed to update song in the database")
	}

	log.Printf("Song with ID %s successfully updated.", req.GetSongId())
	return &proto.UpdateSongResponse{
		Status: "success",
	}, nil
}

func (m *SongService) DeleteSong(ctx context.Context, req *proto.DeleteSongRequest) (*proto.DeleteSongResponse, error) {
	// Prepare the query to delete a song by ID
	query := `DELETE FROM songs WHERE id = $1`

	// Execute the query
	result, err := m.PostgreSQL.ExecContext(ctx, query, req.GetSongId())
	if err != nil {
		log.Printf("Error deleting song with ID %s: %v", req.GetSongId(), err)
		return nil, status.Error(codes.Internal, "Failed to delete song from the database")
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected for deletion: %v", err)
		return nil, status.Error(codes.Internal, "Failed to verify song deletion")
	}
	if rowsAffected == 0 {
		log.Printf("No song found with ID %s to delete.", req.GetSongId())
		return nil, status.Error(codes.NotFound, "No song found with the provided ID")
	}

	log.Printf("Song with ID %s successfully deleted.", req.GetSongId())
	return &proto.DeleteSongResponse{
		Status: "success",
	}, nil
}

func (m *SongService) GetSongLyrics(ctx context.Context, req *proto.GetSongLyricsRequest) (*proto.GetSongLyricsResponse, error) {
	// SQL so'rovi: kupletlarni sahifalash bilan olish
	query := `
        SELECT verse_number, verse_text 
        FROM lyrics 
        WHERE song_id = $1 
        ORDER BY verse_number 
        LIMIT $2 OFFSET $3
    `

	// Sahifalash hisob-kitobi
	pageSize := req.GetPageSize()
	offset := (req.GetPage() - 1) * pageSize

	// SQL so'rovini bajarish
	rows, err := m.PostgreSQL.QueryContext(ctx, query, req.GetSongId(), pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Javobni to'ldirish
	var lyrics []*proto.SongLyrics
	for rows.Next() {
		var lyric proto.SongLyrics
		if err := rows.Scan(&lyric.VerseNumber, &lyric.VerseText); err != nil {
			return nil, err
		}
		lyrics = append(lyrics, &lyric)
	}

	// Javobni qaytarish
	return &proto.GetSongLyricsResponse{
		Lyrics: lyrics,
	}, nil
}
