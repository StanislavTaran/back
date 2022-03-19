DELETE FROM
    `role`
WHERE role IN
       (
        "admin",
       "moderator",
       "user"
       );