package calculator

import (
    "context"
    "encoding/json"
    "time"

    "github.com/redis/go-redis/v9"
)

type LogEntry struct {
    Operation string    `json:"operation"`
    Input     any       `json:"input"`
    Result    any       `json:"result"`
    Time      time.Time `json:"time"`
}

type Repository struct {
    rdb *redis.Client
}

func NewRepository(rdb *redis.Client) *Repository {
    return &Repository{rdb: rdb}
}

func (r *Repository) Save(ctx context.Context, entry LogEntry) error {
    b, err := json.Marshal(entry)
    if err != nil {
        return err
    }
    return r.rdb.RPush(ctx, "calc:logs", b).Err()
}

func (r *Repository) GetRecent(ctx context.Context, count int64) ([]LogEntry, error) {
    vals, err := r.rdb.LRange(ctx, "calc:logs", -count, -1).Result()
    if err != nil {
        return nil, err
    }
    logs := make([]LogEntry, 0, len(vals))
    for _, v := range vals {
        var entry LogEntry
        if err := json.Unmarshal([]byte(v), &entry); err == nil {
            logs = append(logs, entry)
        }
    }
    return logs, nil
}
