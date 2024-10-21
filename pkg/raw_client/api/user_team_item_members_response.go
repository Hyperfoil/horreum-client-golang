package api

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use UserTeamItemMembersGetResponseable instead.
type UserTeamItemMembersResponse struct {
    UserTeamItemMembersGetResponse
}
// NewUserTeamItemMembersResponse instantiates a new UserTeamItemMembersResponse and sets the default values.
func NewUserTeamItemMembersResponse()(*UserTeamItemMembersResponse) {
    m := &UserTeamItemMembersResponse{
        UserTeamItemMembersGetResponse: *NewUserTeamItemMembersGetResponse(),
    }
    return m
}
// CreateUserTeamItemMembersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserTeamItemMembersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserTeamItemMembersResponse(), nil
}
// Deprecated: This class is obsolete. Use UserTeamItemMembersGetResponseable instead.
type UserTeamItemMembersResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserTeamItemMembersGetResponseable
}
