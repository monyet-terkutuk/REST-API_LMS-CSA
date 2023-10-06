-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 07 Okt 2023 pada 00.00
-- Versi server: 10.4.28-MariaDB
-- Versi PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_lms`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `courses`
--

CREATE TABLE `courses` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `title` varchar(225) NOT NULL,
  `materi` text NOT NULL,
  `link_video` text NOT NULL,
  `category` varchar(256) NOT NULL,
  `division` varchar(30) NOT NULL,
  `playlist` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `courses`
--

INSERT INTO `courses` (`id`, `user_id`, `title`, `materi`, `link_video`, `category`, `division`, `playlist`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'mencari jodoh', 'fcvdscfvds', 'dsfsDF', 'BackENd', 'Pemograman', 'Belajar JS', 'qawdeqewd', '2023-08-08 11:41:23', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `course_images`
--

CREATE TABLE `course_images` (
  `id` int(11) NOT NULL,
  `course_id` int(11) NOT NULL,
  `is_primary` tinyint(1) NOT NULL,
  `file_name` varchar(225) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(225) NOT NULL,
  `nim` varchar(255) NOT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(256) DEFAULT NULL,
  `avatar` varchar(225) DEFAULT NULL,
  `role_id` tinyint(1) NOT NULL,
  `division` varchar(20) NOT NULL,
  `no_hp` varchar(15) NOT NULL,
  `alasan_daftar` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `nim`, `email`, `password`, `avatar`, `role_id`, `division`, `no_hp`, `alasan_daftar`, `created_at`, `updated_at`) VALUES
(1, 'Yang Yang', '21110xx1', 'hemm@bukan.com', '$2a$10$RcFO87VMsT5YzhYQEJtiA.uwwqBSkhwKvu1mpxZjXKZf.SgpfzluK', '', 0, 'Multimedia', '00178974', 'Cari Jodoh', '2023-08-08 11:35:32', '2023-08-08 11:35:32'),
(2, 'Boby', '21110xx1', 'hemm@bukan.com', '$2a$10$KcBFAgdUn8drkIYRaf9Qiepsqvdb1snKfXhBC50bBgUmjTs57E13G', 'images/avatar/2-bloodhunt.jpg', 0, 'Multimedia', '00178974', 'Cari Jodoh', '2023-08-08 11:35:57', '2023-08-08 11:36:26');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indeks untuk tabel `course_images`
--
ALTER TABLE `course_images`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `courses`
--
ALTER TABLE `courses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `course_images`
--
ALTER TABLE `course_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `courses`
--
ALTER TABLE `courses`
  ADD CONSTRAINT `courses_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
