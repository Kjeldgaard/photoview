package scanner

import (
	"database/sql"
	"fmt"
	"log"
	"path"

	"github.com/viktorstrate/photoview/api/graphql/models"
	"github.com/viktorstrate/photoview/api/graphql/notification"
)

func ScanPhoto(tx *sql.Tx, photoPath string, albumId int, content_type *string, notificationKey string) error {

	log.Printf("Scanning image: %s\n", photoPath)

	photoName := path.Base(photoPath)

	// Check if image already exists
	row := tx.QueryRow("SELECT (photo_id) FROM photo WHERE path = ?", photoPath)
	var photo_id int64
	if err := row.Scan(&photo_id); err != sql.ErrNoRows {
		if err == nil {
			log.Printf("Image already scanned: %s\n", photoPath)
			return nil
		} else {
			return err
		}
	}

	notifyTimeout := 5000
	notification.BroadcastNotification(&models.Notification{
		Key:     notificationKey,
		Type:    models.NotificationTypeMessage,
		Header:  "Scanning photo",
		Content: fmt.Sprintf("Scanning image at %s", photoPath),
		Timeout: &notifyTimeout,
	})

	result, err := tx.Exec("INSERT INTO photo (title, path, album_id) VALUES (?, ?, ?)", photoName, photoPath, albumId)
	if err != nil {
		log.Printf("ERROR: Could not insert photo into database")
		return err
	}
	photo_id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	row = tx.QueryRow("SELECT * FROM photo WHERE photo_id = ?", photo_id)
	photo, err := models.NewPhotoFromRow(row)
	if err != nil {
		return err
	}

	_, err = ScanEXIF(tx, photo)
	if err != nil {
		log.Printf("ERROR: ScanEXIF for %s: %s\n", photoName, err)
	}

	if err := ProcessPhoto(tx, photo, content_type); err != nil {
		return err
	}

	return nil
}
