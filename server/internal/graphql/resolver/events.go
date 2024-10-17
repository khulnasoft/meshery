package resolver

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/khulnasoft/meshery/server/internal/graphql/model"
	"github.com/khulnasoft/meshery/server/models"
	"github.com/layer5io/meshkit/models/events"
)

func (r *Resolver) eventsResolver(ctx context.Context, provider models.Provider, user models.User) (<-chan *model.Event, error) {
	userID, _ := uuid.FromString(user.ID)
	ch, unsubscribe := r.Config.EventBroadcaster.Subscribe(userID)

	eventsChan := make(chan *model.Event)
	go func(userID uuid.UUID) {
		r.Log.Infof("Events Subscription started for %s", user.ID)
		for {
			select {
			case ech := <-ch:
				event := ech.(*events.Event)
				_event := &model.Event{
					ID:          event.ID.String(),
					UserID:      event.UserID.String(),
					ActedUpon:   event.ActedUpon.String(),
					OperationID: event.OperationID.String(),
					Severity:    model.Severity(event.Severity),
					Description: event.Description,
					Category:    event.Category,
					Action:      event.Action,
					CreatedAt:   event.CreatedAt,
					DeletedAt:   event.DeletedAt,
					UpdatedAt:   event.UpdatedAt,
					Metadata:    event.Metadata,
					Status:      string(event.Status),
					SystemID:    event.SystemID.String(),
				}

				eventsChan <- _event
			case <-ctx.Done():
				unsubscribe()
				close(eventsChan)

				r.Log.Infof("Events Subscription stopped for %s", userID)
				return
			}
		}
	}(userID)
	return eventsChan, nil
}
