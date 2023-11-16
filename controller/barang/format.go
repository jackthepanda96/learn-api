package barang

type BarangRequest struct {
	NamaBarang string `json:"nama_barang"`
	Stok       int    `json:"stok"`
	Harga      int    `json:"harga"`
}
type BarangResponse struct {
	ID         uint   `json:"id"`
	NamaBarang string `json:"nama_barang"`
	Stok       int    `json:"stok"`
	Harga      int    `json:"harga"`
}
