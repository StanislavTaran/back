INSERT INTO users (
       id,
       firstName,
       lastName,
       dataOfBirth,
       email,
       password,
       shortInfo,
       roleId
    )
VALUES (
        "1dc33d06-7f85-4f22-8384-9f3621ed00da",
        "John",
        "Doe",
        CURRENT_TIMESTAMP,
        "johndoe@test.com",
        "password",
        "I am joker",
        3
       );

INSERT INTO edu_institution (
    fullName,
    shortName,
    description
)
VALUES (
    "Harvard College",
        "Harvard",
        "There once was a classical theory Of which quantum disciples were leery. They said, “Why spend so long On a theory that’s wrong?” Well, it works for your everyday query!"
       );

INSERT INTO user_education(
                           userId,
                           eduInstitutionId,
                           eduInstitutionName,
                           faculty,
                           inProgress,
                           startDate,
                           endDate
                           )
VALUES (
        "1dc33d06-7f85-4f22-8384-9f3621ed00da",
        1,
        "",
        "Informsaion technology faculty",
        0,
        CURRENT_TIMESTAMP,
        CURRENT_TIMESTAMP
       );

INSERT INTO company (shortName, fullName, description)
VALUES (
        "Gitlab LTD",
        "Gitlub",
        "Deliver software faster with better security and collaboration in a single platform."
       );

INSERT INTO user_company (
                          userId,
                          companyId,
                          employmentTypeId,
                          companyName,
                          jobTitle,
                          inProgress,
                          startDate,
                          endDate
                          )
VALUES(
       "1dc33d06-7f85-4f22-8384-9f3621ed00da",
       1,
       1,
       "",
       "Full-Stackoferflow developer",
       1,
       CURRENT_TIMESTAMP,
       CURRENT_TIMESTAMP
      );