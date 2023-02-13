// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	gocontext "context"

	"gogs.io/gogs/pkg/conf"
	"gogs.io/gogs/pkg/context"
	"gogs.io/gogs/pkg/db"
	"gogs.io/gogs/pkg/route"
)

const (
	ORGS = "admin/org/list"
)

func Organizations(c *context.Context) {
	c.Data["Title"] = c.Tr("admin.organizations")
	c.Data["PageIsAdmin"] = true
	c.Data["PageIsAdminOrganizations"] = true

	route.RenderUserSearch(c, &route.UserSearchOptions{
		Type: db.UserTypeOrganization,
		Counter: func(gocontext.Context) int64 {
			return db.CountOrganizations()
		},
		Ranger: func(_ gocontext.Context, page, pageSize int) ([]*db.User, error) {
			return db.Organizations(page, pageSize)
		},
		PageSize: conf.UI.Admin.OrgPagingNum,
		OrderBy:  "id ASC",
		TplName:  ORGS,
	})
}
