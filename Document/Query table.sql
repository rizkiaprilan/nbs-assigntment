-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 28, 2021 at 03:52 PM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `nbs`
--

-- --------------------------------------------------------

--
-- Table structure for table `employees`
--

CREATE TABLE `employees` (
  `id` varchar(200) NOT NULL,
  `fullname` varchar(200) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `isAdmin` tinyint(1) NOT NULL DEFAULT 0,
  `verified` tinyint(1) NOT NULL DEFAULT 0,
  `created_at` date NOT NULL DEFAULT current_timestamp(),
  `updated_at` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `employees`
--

INSERT INTO `employees` (`id`, `fullname`, `email`, `password`, `phone`, `isAdmin`, `verified`, `created_at`, `updated_at`) VALUES
('d29c51a3-7abb-4399-a360-32043b153f3c', 'rizki aprilan', 'riskiazza@gmail.com', '$2a$12$BoFHY/npd0xPyVKd4rtk8OLPF.LJEjqrWWLW97Bsn0OYdRDmbZeiC', '082170725072', 1, 1, '2021-03-27', '2021-03-27');

-- --------------------------------------------------------

--
-- Table structure for table `employees_performance`
--

CREATE TABLE `employees_performance` (
  `id` varchar(200) NOT NULL,
  `employee_id` varchar(200) NOT NULL,
  `score` bigint(20) NOT NULL DEFAULT 0,
  `created_at` date NOT NULL DEFAULT current_timestamp(),
  `updated_at` date NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `employees_performance`
--

INSERT INTO `employees_performance` (`id`, `employee_id`, `score`, `created_at`, `updated_at`) VALUES
('b21b8542-97b1-4df0-9cec-2fe96199b33d', 'd29c51a3-7abb-4399-a360-32043b153f3c', 7, '2021-03-28', '2021-03-28');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `employees`
--
ALTER TABLE `employees`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `employees_performance`
--
ALTER TABLE `employees_performance`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_employee_id` (`employee_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `employees_performance`
--
ALTER TABLE `employees_performance`
  ADD CONSTRAINT `fk_employee_id` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
