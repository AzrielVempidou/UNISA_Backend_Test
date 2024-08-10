package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Presensi represents the structure of a student attendance record in MongoDB
type Presensi struct {
    ID              primitive.ObjectID `bson:"_id,omitempty"`
    NamaLengkap     string             `bson:"nama_lengkap"`
    Alamat          string             `bson:"alamat"`
    TanggalLahir    string             `bson:"tanggal_lahir"`
    Email           string             `bson:"email"`
    NoHp            int             `bson:"no_hp"`
    NamaPTInstansi  string             `bson:"nama_pt_instansi"`
    Jabatan         string             `bson:"jabatan"`
    NamaInstagram   string             `bson:"nama_instagram,omitempty"`
    NamaFacebook    string             `bson:"nama_facebook,omitempty"`
    TanggalPresensi string             `bson:"tanggal_presensi"`
    Status          string             `bson:"status"`
}
