-- name: CreateFeedback :exec
INSERT INTO "feedback" (
        id,
        user_id,
        type,
        message,
        solved,
        created_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        false,
        now()
    );

/* SOLVING FEEDBACK QUERY WILL BE ADDED
 -- name: SolveFeedback :exec
 */