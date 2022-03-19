DELETE FROM
            `employment_type`
WHERE type IN
    (
    "full-time",
    "part-time",
    "self-employed",
    "freelance",
    "contract",
    "intership",
    "seasonal"
    );
