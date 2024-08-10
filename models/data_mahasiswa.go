// models/model.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Mahasiswa represents the structure of a student record in MongoDB
type Mahasiswa struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    NamaLengkap     string             `bson:"nama_lengkap"`
    Alamat          string             `bson:"alamat"`
    TempatLahir     string             `bson:"tempat_lahir"`
    TanggalLahir    string             `bson:"tanggal_lahir"`
    Email           string             `bson:"email"`
    NoHp            int             `bson:"no_hp"`
    NamaPTInstansi  string             `bson:"nama_pt_instansi"`
    Jabatan         string             `bson:"jabatan"`
    NamaInstagram   string             `bson:"nama_instagram,omitempty"`
    NamaFacebook    string             `bson:"nama_facebook,omitempty"`
    SumberInformasi string             `bson:"sumber_informasi,omitempty"` // Only for registration
}
