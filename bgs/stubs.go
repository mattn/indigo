package bgs

import (
	"io"

	comatprototypes "github.com/bluesky-social/indigo/api/atproto"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
)

func (s *BGS) RegisterHandlersAppBsky(e *echo.Echo) error {
	return nil
}

func (s *BGS) RegisterHandlersComAtproto(e *echo.Echo) error {
	e.GET("/xrpc/com.atproto.sync.getCheckout", s.HandleComAtprotoSyncGetCheckout)
	e.GET("/xrpc/com.atproto.sync.getCommitPath", s.HandleComAtprotoSyncGetCommitPath)
	e.GET("/xrpc/com.atproto.sync.getHead", s.HandleComAtprotoSyncGetHead)
	e.GET("/xrpc/com.atproto.sync.getRecord", s.HandleComAtprotoSyncGetRecord)
	e.GET("/xrpc/com.atproto.sync.getRepo", s.HandleComAtprotoSyncGetRepo)
	return nil
}

func (s *BGS) HandleComAtprotoSyncGetCheckout(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetCheckout")
	defer span.End()
	commit := c.QueryParam("commit")
	did := c.QueryParam("did")
	var out io.Reader
	var handleErr error
	// func (s *BGS) handleComAtprotoSyncGetCheckout(ctx context.Context,commit string,did string) (io.Reader, error)
	out, handleErr = s.handleComAtprotoSyncGetCheckout(ctx, commit, did)
	if handleErr != nil {
		return handleErr
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *BGS) HandleComAtprotoSyncGetCommitPath(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetCommitPath")
	defer span.End()
	did := c.QueryParam("did")
	earliest := c.QueryParam("earliest")
	latest := c.QueryParam("latest")
	var out *comatprototypes.SyncGetCommitPath_Output
	var handleErr error
	// func (s *BGS) handleComAtprotoSyncGetCommitPath(ctx context.Context,did string,earliest string,latest string) (*comatprototypes.SyncGetCommitPath_Output, error)
	out, handleErr = s.handleComAtprotoSyncGetCommitPath(ctx, did, earliest, latest)
	if handleErr != nil {
		return handleErr
	}
	return c.JSON(200, out)
}

func (s *BGS) HandleComAtprotoSyncGetHead(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetHead")
	defer span.End()
	did := c.QueryParam("did")
	var out *comatprototypes.SyncGetHead_Output
	var handleErr error
	// func (s *BGS) handleComAtprotoSyncGetHead(ctx context.Context,did string) (*comatprototypes.SyncGetHead_Output, error)
	out, handleErr = s.handleComAtprotoSyncGetHead(ctx, did)
	if handleErr != nil {
		return handleErr
	}
	return c.JSON(200, out)
}

func (s *BGS) HandleComAtprotoSyncGetRecord(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetRecord")
	defer span.End()
	collection := c.QueryParam("collection")
	commit := c.QueryParam("commit")
	did := c.QueryParam("did")
	rkey := c.QueryParam("rkey")
	var out io.Reader
	var handleErr error
	// func (s *BGS) handleComAtprotoSyncGetRecord(ctx context.Context,collection string,commit string,did string,rkey string) (io.Reader, error)
	out, handleErr = s.handleComAtprotoSyncGetRecord(ctx, collection, commit, did, rkey)
	if handleErr != nil {
		return handleErr
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}

func (s *BGS) HandleComAtprotoSyncGetRepo(c echo.Context) error {
	ctx, span := otel.Tracer("server").Start(c.Request().Context(), "HandleComAtprotoSyncGetRepo")
	defer span.End()
	did := c.QueryParam("did")
	earliest := c.QueryParam("earliest")
	latest := c.QueryParam("latest")
	var out io.Reader
	var handleErr error
	// func (s *BGS) handleComAtprotoSyncGetRepo(ctx context.Context,did string,earliest string,latest string) (io.Reader, error)
	out, handleErr = s.handleComAtprotoSyncGetRepo(ctx, did, earliest, latest)
	if handleErr != nil {
		return handleErr
	}
	return c.Stream(200, "application/vnd.ipld.car", out)
}
