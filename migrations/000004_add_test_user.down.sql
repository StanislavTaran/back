SET FOREIGN_KEY_CHECKS=0; -- to disable them

DELETE FROM `users` WHERE id IN (
        "1dc33d06-7f85-4f22-8384-9f3621ed00da"
        );


DELETE FROM `user_company` WHERE id = 1;
DELETE FROM `company` WHERE id = 1;

DELETE FROM `user_education`  WHERE id = 1;
DELETE FROM `edu_institution`  WHERE id = 1;

SET FOREIGN_KEY_CHECKS=1; -- to re-enable them