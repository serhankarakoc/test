package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete için GORM'un desteklediği alan
	DeletedBy *uint
}

// Kaydedilmeden önce CreatedAt ve CreatedBy için varsayılan değerleri ayarlayın
func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if bm.CreatedAt.IsZero() {
		bm.CreatedAt = time.Now()
	}
	if bm.CreatedBy == 0 {
		bm.CreatedBy = getCurrentUserID()
	}
	return nil
}

// Güncellenmeden önce UpdatedAt ve UpdatedBy için varsayılan değerleri ayarlayın
func (bm *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	bm.UpdatedAt = time.Now()
	if bm.UpdatedBy == 0 {
		bm.UpdatedBy = getCurrentUserID()
	}
	return nil
}

// Silinmeden önce DeletedBy'ı ayarlayın
func (bm *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	userID := getCurrentUserID()
	bm.DeletedBy = &userID
	bm.DeletedAt.Time = time.Now()
	bm.DeletedAt.Valid = true

	// Soft delete'i manuel olarak işaretleyin
	err = tx.Model(bm).UpdateColumns(map[string]interface{}{
		"deleted_at": bm.DeletedAt,
		"deleted_by": bm.DeletedBy,
	}).Error
	if err != nil {
		return err
	}
	// Asıl silme işlemini durdur
	return gorm.ErrInvalidData
}

// Geçerli kullanıcı kimliğini almak için dummy (yalancı) fonksiyon
// Gerçek bir uygulamada, bu muhtemelen daha fazla mantık veya bağlam gerektirir
func getCurrentUserID() uint {
	return 1 // Geçerli kullanıcı kimliğini almak için gerçek mantıkla değiştirin
}
