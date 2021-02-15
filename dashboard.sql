-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Feb 15, 2021 at 10:32 AM
-- Server version: 10.1.38-MariaDB
-- PHP Version: 7.1.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `loena_dashboard`
--

-- --------------------------------------------------------

--
-- Table structure for table `adhoc`
--

CREATE TABLE `adhoc` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `filename` varchar(255) NOT NULL,
  `status` varchar(10) NOT NULL DEFAULT 'SUBMIT',
  `scheduled_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `adhoc`
--

INSERT INTO `adhoc` (`id`, `name`, `filename`, `status`, `scheduled_at`, `created_at`, `updated_at`) VALUES
(2, 'program loan april', 'logger.txt', 'SUBMIT', '2020-03-01 00:00:00', '2021-02-15 10:11:04', '2021-02-15 10:11:04'),
(3, 'program loan april', 'logger.txt', 'SUBMIT', '2020-03-01 00:00:00', '2021-02-15 10:15:15', '2021-02-15 10:15:15');

-- --------------------------------------------------------

--
-- Table structure for table `menus`
--

CREATE TABLE `menus` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `icon` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `sorting` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `menus`
--

INSERT INTO `menus` (`id`, `name`, `path`, `icon`, `parent_id`, `is_active`, `sorting`, `created_at`, `updated_at`, `created_by`) VALUES
(1, 'User Management', '/loena/v1/users', 'fa-user', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL),
(2, 'Menu Management', '/loena/v1/menus', 'fa-list', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL),
(3, 'Customer', '/loena/v1/customer', 'fa-person', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL),
(4, 'History', '/loena/v1/history', 'fa-history', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL),
(5, 'Reminder List', '/loena/v1/reminder', 'fa-calendar', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL),
(6, 'Adhoc', '/loena/v1/adhoc', 'fa-upload', 0, 1, NULL, '2021-02-15 03:24:56', '2021-02-15 03:24:56', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `menus_roles`
--

CREATE TABLE `menus_roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `menus_id` int(11) DEFAULT NULL,
  `roles_id` int(11) DEFAULT NULL,
  `is_create` tinyint(4) NOT NULL DEFAULT '0',
  `is_read` tinyint(4) NOT NULL DEFAULT '1',
  `is_edit` tinyint(4) NOT NULL DEFAULT '0',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `menus_roles`
--

INSERT INTO `menus_roles` (`id`, `menus_id`, `roles_id`, `is_create`, `is_read`, `is_edit`, `is_delete`, `created_at`, `updated_at`, `created_by`) VALUES
(1, 1, 1, 1, 1, 1, 1, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(2, 2, 1, 1, 1, 1, 1, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(3, 6, 1, 1, 1, 1, 1, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(4, 3, 1, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(5, 4, 1, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(6, 5, 1, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(7, 6, 2, 1, 1, 1, 1, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(8, 3, 2, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(9, 4, 2, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(10, 5, 2, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(11, 6, 3, 1, 1, 1, 1, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(12, 3, 3, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(13, 4, 3, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL),
(14, 5, 3, 0, 1, 0, 0, '2021-02-15 10:29:37', '2021-02-15 10:29:37', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_superadmin` tinyint(1) DEFAULT NULL,
  `theme_color` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`, `is_superadmin`, `theme_color`, `created_at`, `updated_at`, `created_by`) VALUES
(1, 'Super Administrator', 1, 'skin-green-light', '2020-04-22 08:52:21', '2020-04-22 08:52:21', NULL),
(2, 'Admin', 0, 'skin-green-light', '2020-04-22 08:52:21', '2020-04-22 08:52:21', NULL),
(3, 'User', 0, 'skin-green-light', '2020-04-22 08:52:21', '2020-04-22 08:52:21', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `settings`
--

CREATE TABLE `settings` (
  `id` int(11) NOT NULL,
  `key` varchar(100) NOT NULL DEFAULT '',
  `value` varchar(100) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `photo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `roles_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `created_by` int(11) DEFAULT NULL,
  `status` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'active'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `photo`, `email`, `password`, `roles_id`, `created_at`, `updated_at`, `created_by`, `status`) VALUES
(1, 'super admin', '', 'superadmin@tsel.co.id', '76419c58730d9f35de7ac538c2fd6737', 1, '2019-07-23 18:50:58', '2020-07-11 23:00:46', NULL, 'active'),
(2, 'admin', NULL, 'admin@tsel.co.id', '76419c58730d9f35de7ac538c2fd6737', 2, '2019-07-24 08:53:41', '2020-07-11 23:00:37', NULL, 'active'),
(3, 'user', NULL, 'user@tsel.co.id', '76419c58730d9f35de7ac538c2fd6737', 3, '2019-07-24 08:54:08', '2020-07-11 23:00:24', NULL, 'active');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `adhoc`
--
ALTER TABLE `adhoc`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `menus_roles`
--
ALTER TABLE `menus_roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `settings`
--
ALTER TABLE `settings`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `adhoc`
--
ALTER TABLE `adhoc`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `menus`
--
ALTER TABLE `menus`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `menus_roles`
--
ALTER TABLE `menus_roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `settings`
--
ALTER TABLE `settings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
