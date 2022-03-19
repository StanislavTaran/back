CREATE TABLE `users` (
                         `id` varchar(255) PRIMARY KEY,
                         `firstName` varchar(255),
                         `lastName` varchar(255),
                         `dataOfBirth` datetime,
                         `email` varchar(255),
                         `password` varchar(255),
                         `shortInfo` varchar(255),
                         `roleId` INT,
                         `createdAt` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `updatedAt` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `edu_institution` (
                                  `id` int AUTO_INCREMENT PRIMARY KEY,
                                  `shortName` varchar(255),
                                  `fullName` varchar(255),
                                  `description` text
);

CREATE TABLE `employment_type` (
                                   `id` int AUTO_INCREMENT PRIMARY KEY,
                                   `type` varchar(255)
);

CREATE TABLE `user_education` (
                                  `id` int AUTO_INCREMENT PRIMARY KEY,
                                  `userId` varchar(255),
                                  `eduInstitutionId` int,
                                  `eduInstitutionName` varchar(255),
                                  `faculty` varchar(255),
                                  `inProgress` tinyint,
                                  `startDate` timestamp,
                                  `endDate` timestamp
);

CREATE TABLE `company` (
                           `id` int AUTO_INCREMENT PRIMARY KEY,
                           `fullName` varchar(255),
                           `shortName` varchar(255),
                           `description` text
);

CREATE TABLE `user_company` (
                                `id` int AUTO_INCREMENT PRIMARY KEY,
                                `userId` varchar(255),
                                `companyId` int,
                                `employmentTypeId` int,
                                `companyName` varchar(255),
                                `jobTitle` varchar(255),
                                `inProgress` tinyint,
                                `startDate` timestamp,
                                `endDate` timestamp
);

CREATE TABLE `role` (
                                `id` int AUTO_INCREMENT PRIMARY KEY,
                                `role` varchar(255)
);

ALTER TABLE company AUTO_INCREMENT=1;
ALTER TABLE user_company AUTO_INCREMENT=1;
ALTER TABLE user_company AUTO_INCREMENT=1;
ALTER TABLE edu_institution AUTO_INCREMENT=1;
ALTER TABLE employment_type AUTO_INCREMENT=1;
ALTER TABLE role AUTO_INCREMENT=1;

ALTER TABLE `user_education` ADD FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `user_education` ADD FOREIGN KEY (`eduInstitutionId`) REFERENCES `edu_institution` (`id`);

ALTER TABLE `user_company` ADD FOREIGN KEY (`companyId`) REFERENCES `company` (`id`);

ALTER TABLE `user_company` ADD FOREIGN KEY (`employmentTypeId`) REFERENCES `employment_type` (`id`);

ALTER TABLE `user_company` ADD FOREIGN KEY (`userId`) REFERENCES `users` (`id`);

ALTER TABLE `users` ADD FOREIGN KEY (`roleId`) REFERENCES `role` (`id`);
