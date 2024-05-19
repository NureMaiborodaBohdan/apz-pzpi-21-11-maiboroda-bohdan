CREATE TABLE `Location` (
                            `LocationID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                            `Country` VARCHAR(255) NOT NULL,
                            `City` VARCHAR(255) NOT NULL,
                            `Adress` VARCHAR(255) NOT NULL,
                            `PostCode` INT NOT NULL
);

CREATE TABLE `ThresholdValues` (
                                   `ThresholdID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                   `LegalLimit` DOUBLE NOT NULL
);

CREATE TABLE `Company` (
                           `CompanyID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                           `Name` VARCHAR(255) NOT NULL,
                           `Description` TEXT NOT NULL,
                           `LocationID` INT UNSIGNED NOT NULL,
                           `ThresholdID` INT UNSIGNED NOT NULL,
                           FOREIGN KEY (`LocationID`) REFERENCES `Location` (`LocationID`) ON DELETE CASCADE,
                           FOREIGN KEY (`ThresholdID`) REFERENCES `ThresholdValues` (`ThresholdID`) ON DELETE CASCADE
);

CREATE TABLE `User` (
                        `UserID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        `Username` VARCHAR(255) NOT NULL,
                        `Email` VARCHAR(255) NOT NULL,
                        `Password` VARCHAR(255) NOT NULL,
                        `Role` ENUM('Admin', 'User') NOT NULL,
                        `Name` VARCHAR(255) NOT NULL,
                        `Surname` VARCHAR(255) NOT NULL,
                        `Patronymic` VARCHAR(255) NOT NULL,
                        `CompanyID` INT UNSIGNED NOT NULL,
                        `Sex` ENUM('Male', 'Female') NOT NULL,
                        FOREIGN KEY (`CompanyID`) REFERENCES `Company` (`CompanyID`) ON DELETE CASCADE
);

CREATE TABLE `TestResult` (
                              `TestID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                              `UserID` INT UNSIGNED NOT NULL,
                              `TestTime` DATETIME NOT NULL,
                              `AlcoholLevel` DOUBLE NOT NULL,
                              `IsDrunk` BOOL NOT NULL,
                              `Description` TEXT NOT NULL,
                              FOREIGN KEY (`UserID`) REFERENCES `User` (`UserID`) ON DELETE CASCADE
);

CREATE TABLE `Notification` (
                                `NotificationID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                `Message` TEXT NOT NULL,
                                `SentTime` DATETIME NOT NULL,
                                `UserID` INT UNSIGNED NOT NULL,
                                FOREIGN KEY (`UserID`) REFERENCES `User` (`UserID`) ON DELETE CASCADE
);

CREATE TABLE `AccessControl` (
                                 `AccessID` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                 `UserID` INT UNSIGNED NOT NULL,
                                 `AccessTime` DATETIME NOT NULL,
                                 FOREIGN KEY (`UserID`) REFERENCES `User` (`UserID`) ON DELETE CASCADE
);