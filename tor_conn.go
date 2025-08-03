package tor_conn

import (
    "context"
    "log"
    "time"

    "github.com/cretz/bine/tor"
)

func StartTor() (*tor.Tor, string, error) {
    t, err := tor.Start(nil, nil)
    if err != nil {
        return nil, "", err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := t.WaitForNetwork(ctx); err != nil {
        return nil, "", err
    }

    return t, t.ProxyNetListener().Addr().String(), nil
}
