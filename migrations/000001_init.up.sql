CREATE TABLE `users` (
                         `id` varchar(255) PRIMARY KEY UNIQUE,
                         `firstName` varchar(255),
                         `lastName` varchar(255),
                         `email` varchar(255) UNIQUE,
                         `createdAt` timestamp DEFAULT CURRENT_TIMESTAMP,
                         `updatedAt` timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `countries` (
                             `code` int PRIMARY KEY UNIQUE,
                             `shortName` varchar(255) UNIQUE,
                             `fullName` varchar(255) UNIQUE
);

CREATE TABLE `regions` (
                           `code` int PRIMARY KEY UNIQUE,
                           `countryCode` int,
                           `shortName` varchar(255) UNIQUE,
                           `fullName` varchar(255) UNIQUE
);

INSERT INTO `countries` (code, shortName, fullName) VALUES (1, "UA", "Ukraine");
INSERT INTO `regions` (code, countryCode, shortName, fullName) VALUES (1,1,"KH", "Kharkiv");
INSERT INTO `users` (id, firstName, lastName, email) VALUES ("test","test","test","test@test.com");
