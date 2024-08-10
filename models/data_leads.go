package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataLeads struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NamaLengkap string             `bson:"nama_lengkap"`
	Alamat      string             `bson:"alamat"`
	TempatLahir string             `bson:"tempat_lahir"`
	TanggalLahir string            `bson:"tanggal_lahir"`
	Email       string             `bson:"email"`
	NoHp        string             `bson:"no_hp"`
	NamaPTInstansi string          `bson:"nama_pt_instansi"`
	Jabatan     string             `bson:"jabatan"`
}
