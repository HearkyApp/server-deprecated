/*
 * Hearky Server
 * Copyright (C) 2021 Hearky
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"github.com/gofiber/websocket/v2"
	"go.uber.org/zap"
)

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

func login(c *websocket.Conn, m *WSMessage) {
	zap.L().Info("mes", zap.Any("message", m.Data))
	req, ok := m.Data.(*LoginRequest)
	if !ok {
		sendError(c, ErrInvalidPayload)
		return
	}
	zap.L().Info("Login Request", zap.Any("req", req))
}
