# Rumah Makan - PT Qoin Digital Indonesia
Sistem ini dirancang sebagai bagian dari proses rekruitment di PT Qoin Digital Indonesia tahap technical test.

## Background
Rancangan sistem ini dimaksudkan untuk mengoptimalkan proses bisnis yang ada di rumah makan. 
<br />
Kebutuhan utama dari sistem yang harus dipenuhi :
1. Aplikasi ini bisa memasukkan pesanan-pesanan makanan pelanggan
2. Aplikasi ini bisa mengeluarkan struk pembelian
3. Aplikasi ini bisa mengeluarkan laporan penghasilan mingguan dan bulanan
4. Aplikasi ini bisa mengeluarkan laporan stok

Batasan :
1. Hanya waiters yang dapat memasukkan pesanan pelanggan dan mengajukan struk pembelian kepada kasir jika pelanggan telah selesai
2. Pesanan akan di kirim ke kitchen atau barista tergantung menu yang dipesan
3. Struk pembelian akan dikeluarkan oleh kasir setelah mendapat pengajuan dari waiters, kemudian setelah pembayaran diterima kasir akan mengklik tombol konfirmasi
4. Laporan penghasilan dan stok hanya bisa di cetak oleh admin

Entitas :
1. Admin :
- Menambahkan dan menghapus data karyawan serta memberikan role
- Menambah, merubah, atau menghapus data menu
- Mengeluarkan laporan penghasilan dan melihat history transaksi
- Mengeluarkan laporan stok

2. Kasir
- Menerima pembayaran
- Mencetak struk pembelian

3. Waiters
- Meng-input pesanan pelanggan
- Konfirmasi jika pesanan sudah di antar

4. Kitchen
- Membuat pesanan
- Konfirmasi pesanan setelah selesai dibuat

5. Barista
- Membuat pesanan
- Konfirmasi pesanan setelah selesai dibuat

## Desain Arsitektur

## Desain Database

## API Specification