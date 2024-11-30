package handler

import (
	"apigateway/proto/musicproto"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
)

type SongClient struct {
	Client musicproto.SongServiceClient
}

func NewSongClient(cl musicproto.SongServiceClient) *SongClient {
	return &SongClient{Client: cl}
}

// AddSong godoc
// @Router		/api/songs/add [post]
// @Summary 	Add a song
// @Description Adds new song information
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		song body musicproto.AddSongRequest true "Song information"
// @Success 	200 {object} musicproto.AddSongResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) AddSong(c *fiber.Ctx) error {
	var addSongReq musicproto.AddSongRequest

	// Reading the Request Body
	if err := protojson.Unmarshal(c.Body(), &addSongReq); err != nil {
		log.Println("AddSong - Unmarshal error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "AddSong - Error reading the data",
		})
	}

	// Sending AddSong request via gRPC
	resp, err := s.Client.AddSong(c.Context(), &addSongReq)
	if err != nil {
		log.Println("s.Client.AddSong:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetSong godoc
// @Router		/api/songs/id [get]
// @Summary 	Get a song by ID
// @Description Retrieves a song by its ID
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		song_id query string true "Song ID"
// @Success 	200 {object} musicproto.GetSongResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	404 {object} models.StandartError "Not found error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) GetSong(c *fiber.Ctx) error {
	songID := c.Query("song_id") // Get song_id from URL parameter

	// Sending GetSong request via gRPC
	resp, err := s.Client.GetSong(c.Context(), &musicproto.GetSongRequest{SongId: songID})
	if err != nil {
		log.Println("s.Client.GetSong:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetSongs godoc
// @Router		/api/songs/all [get]
// @Summary 	Get songs (with filtering and pagination)
// @Description Retrieves songs with filtering and pagination
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		group query string false "Group name"
// @Param 		page query int false "Page number"
// @Param 		page_size query int false "Number of items per page"
// @Success 	200 {object} musicproto.GetSongsResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) GetSongs(c *fiber.Ctx) error {
	group := c.Query("group")
	page := c.QueryInt("page", 1)           // Default page=1
	pageSize := c.QueryInt("page_size", 10) // Default page size=10

	// Sending GetSongs request via gRPC
	resp, err := s.Client.GetSongs(c.Context(), &musicproto.GetSongsRequest{
		Group:    group,
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	if err != nil {
		log.Println("s.Client.GetSongs:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}

// UpdateSong godoc
// @Router		/api/songs/update [put]
// @Summary 	Update song information
// @Description Updates the song's information
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		song body musicproto.UpdateSongRequest true "Updated song information"
// @Success 	200 {object} musicproto.UpdateSongResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) UpdateSong(c *fiber.Ctx) error {
	var updateSongReq musicproto.UpdateSongRequest

	// Reading the Request Body
	if err := protojson.Unmarshal(c.Body(), &updateSongReq); err != nil {
		log.Println("UpdateSong - Unmarshal error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UpdateSong - Error reading the data",
		})
	}

	// Sending UpdateSong request via gRPC
	resp, err := s.Client.UpdateSong(c.Context(), &updateSongReq)
	if err != nil {
		log.Println("s.Client.UpdateSong:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}

// DeleteSong godoc
// @Router		/api/songs/delete [delete]
// @Summary 	Delete a song
// @Description Deletes a song
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		song_id query string true "Song ID"
// @Success 	200 {object} musicproto.DeleteSongResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) DeleteSong(c *fiber.Ctx) error {
	songID := c.Query("song_id")

	// Sending DeleteSong request via gRPC
	resp, err := s.Client.DeleteSong(c.Context(), &musicproto.DeleteSongRequest{SongId: songID})
	if err != nil {
		log.Println("s.Client.DeleteSong:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}

// GetSongLyrics godoc
// @Router		/api/songs/lyrics/id [get]
// @Summary 	Get song lyrics
// @Description Retrieves song lyrics by ID
// @Tags 		Songs
// @Accept 		json
// @Produce 	json
// @Param 		song_id query string true "Song ID"
// @Success 	200 {object} musicproto.GetSongLyricsResponse
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	404 {object} models.StandartError "Not found error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (s *SongClient) GetSongLyrics(c *fiber.Ctx) error {
	// Getting song_id from URL parameter
	songID := c.Query("song_id")

	// Sending GetSongLyrics request via gRPC
	resp, err := s.Client.GetSongLyrics(c.Context(), &musicproto.GetSongLyricsRequest{SongId: songID})
	if err != nil {
		log.Println("s.Client.GetSongLyrics:", err)
		// Returning 404 error if song not found
		if err.Error() == "not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Song lyrics not found",
			})
		}
		// Returning 500 error for other errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Sending response in JSON format
	return c.Status(fiber.StatusOK).JSON(resp)
}
