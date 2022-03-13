CREATE TABLE `users` (
                         `id` varchar(255) PRIMARY KEY UNIQUE,
                         `firstName` varchar(255),
                         `lastName` varchar(255),
                         `countryCode` int,
                         `regionCode` int,
                         `createdAt` timestamp,
                         `updatedAt` timestamp
);

CREATE TABLE `countries` (
                             `code` int PRIMARY KEY UNIQUE,
                             `shortName` varchar(255),
                             `fullName` varchar(255)
);

CREATE TABLE `regions` (
                           `code` int PRIMARY KEY UNIQUE,
                           `countryCode` int,
                           `shortName` varchar(255),
                           `fullName` varchar(255)
);

ALTER TABLE `users` ADD FOREIGN KEY (`countryCode`) REFERENCES `countries` (`code`);

ALTER TABLE `regions` ADD FOREIGN KEY (`countryCode`) REFERENCES `countries` (`code`);
