package gomaildriver

import (
    "time"
)

type MailDriver interface {
    SetPageSize(int) error
    SetNumberOfPages(int) error
    SetPageOffset(int) error
    SetLatest(time.Time) error
    SetEarliest(time.Time) error
    Fetch() ([]string, error)
}
