package consulv

import (
	"context"
	"log/slog"
	"testing"

	"github.com/abmpio/configurationx/options/consul"
	"github.com/stretchr/testify/assert"
)

func TestReadFromConsul_DefaultLoggerEnablesDebug(t *testing.T) {
	c := ReadFromConsul(consul.ConsulOptions{}, nil)
	assert.True(t, c.Logger.Enabled(context.Background(), slog.LevelDebug))
}

