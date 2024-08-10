package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ListTautan struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	NamaProgram       string             `bson:"nama_program"`
	NamaInstansi      string             `bson:"nama_instansi"`
	NamaKegiatan      string             `bson:"nama_kegiatan"`
	Alamat            string             `bson:"alamat"`
	NamaPIC           string             `bson:"nama_pic"`
	NamaPICPTInstansi string             `bson:"nama_pic_pt_instansi"`
	TanggalMulai      string             `bson:"tanggal_mulai"`
	TanggalAkhir      string             `bson:"tanggal_akhir"`
	QRCodePath        string             `bson:"qr_code_path"` // New field for QR code path
}
