# About Task Tracker CLI

## Visi Proyek (The "Why")
Task Tracker CLI adalah aplikasi *command-line* yang dirancang untuk mengelola pekerjaan sehari-hari dengan overhead minimum. Tidak ada *database server* yang berat, tidak ada *login*, hanya eksekusi dan pencatatan yang cepat via terminal. Aplikasi ini cocok untuk pengembang atau *power-user* yang menghabiskan sebagian besar waktunya di terminal.

## Arsitektur & Teknologi Utama
Aplikasi ini dikembangkan dengan bahasa pemrograman **Go (Golang)** yang mengutamakan kecepatan eksekusi dan *single binary output*. 

Secara konsep arsitektur, aplikasi ini mengadopsi prinsip **Clean Architecture** sederhana:
1. **Presentation Layer (`cmd/`)**: Menggunakan dependensi *cobra* (atau serupa) untuk menangani eksekusi perintah terminal (CLI routing) seperti `add`, `list`, `update`, `delete`.
2. **Business/Domain Layer (`internal/task/`)**: Berisi inti logika dari pelacakan tugas (memvalidasi tugas, mengubah status). Komponen ini tidak peduli apakah data disimpan di JSON atau PostgreSQL.
3. **Infrastructure/Data Layer (`pkg/`)**: Menangani interaksi ke *file system*. Semua operasi tulis/baca dialirkan menjadi struktur data Go dari dan ke file `tasks.json`.

## Filosofi Desain (Pareto 80/20)
- **Fokus pada 20% Fitur Utama**: Menambahkan (`add`) dan melihat daftar tugas (`list`) adalah 80% dari apa yang dibutuhkan pengguna pengelola tugas. Fitur ini dioptimasi agar tanpa hambatan.
- **Penyimpanan Lokal Sederhana**: Menggunakan representasi **JSON file** murni. Menghilangkan kebutuhan migrasi *database* (SQL/NoSQL) di tahap awal pengembangan, karena tidak menyelesaikan masalah utama (*managing daily to-dos*).
- **Extensible**: Logika penyimpanan sengaja dipisahkan di `pkg/fs.go`. Jika aplikasi ini *scale up* dan pengguna membutuhkan sinkronisasi via *cloud backend* atau database SQLite, pergeseran hanya perlu dilakukan di *layer data*, tanpa menyentuh *domain logic*.

## Rencana Pengembangan Lanjut (Roadmap)
Beberapa inisiatif arsitektural yang dapat ditambahkan di fase berikutnya jika kebutuhan bertambah:
- **Concurrency**: Integrasi *goroutine* jika di masa depan CLI harus melakukan *fetching/sync* ke API *remote*.
- **Observability**: Menambahkan *structured logging* (slog) untuk *debugging* proses yang gagal (terutama terkait *I/O errors* saat baca/tulis JSON).
- **Testing**: Menerapkan *Table-Driven Tests* khas ekosistem Go untuk memastikan *domain logic* lulus semua *edge cases* (tugas kosong, ID tidak valid, *handling* file JSON korup).
