// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package keychain

import "fmt"

// UserPasswd returns the user name and password for authenticating
// to the named server. If the user argument is non-empty, UserPasswd
// restricts its search to passwords for the named user.
func UserPasswd(server, preferredUser string) (string, string, error) {
	user, passwd, err := userPasswd(server, preferredUser)
	if err != nil {
		if preferredUser != "" {
			return "", "", fmt.Errorf("error: loading password for %s@%s: %v", preferredUser, server, err)
		}
		return "", "", fmt.Errorf("error: loading password for %s: %v", server, err)
	}

	return user, passwd, nil
}
