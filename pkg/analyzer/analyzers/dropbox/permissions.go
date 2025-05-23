// Code generated by go generate; DO NOT EDIT.
package dropbox

import "errors"

type Permission int

const (
    Invalid Permission = iota
    AccountInfoWrite Permission = iota
    AccountInfoRead Permission = iota
    FilesMetadataWrite Permission = iota
    FilesMetadataRead Permission = iota
    FilesContentWrite Permission = iota
    FilesContentRead Permission = iota
    SharingWrite Permission = iota
    SharingRead Permission = iota
    FileRequestsWrite Permission = iota
    FileRequestsRead Permission = iota
    ContactsWrite Permission = iota
    ContactsRead Permission = iota
    Openid Permission = iota
    Profile Permission = iota
    Email Permission = iota
)

var (
    PermissionStrings = map[Permission]string{
        AccountInfoWrite: "account_info.write",
        AccountInfoRead: "account_info.read",
        FilesMetadataWrite: "files.metadata.write",
        FilesMetadataRead: "files.metadata.read",
        FilesContentWrite: "files.content.write",
        FilesContentRead: "files.content.read",
        SharingWrite: "sharing.write",
        SharingRead: "sharing.read",
        FileRequestsWrite: "file_requests.write",
        FileRequestsRead: "file_requests.read",
        ContactsWrite: "contacts.write",
        ContactsRead: "contacts.read",
        Openid: "openid",
        Profile: "profile",
        Email: "email",
    }

    StringToPermission = map[string]Permission{
        "account_info.write": AccountInfoWrite,
        "account_info.read": AccountInfoRead,
        "files.metadata.write": FilesMetadataWrite,
        "files.metadata.read": FilesMetadataRead,
        "files.content.write": FilesContentWrite,
        "files.content.read": FilesContentRead,
        "sharing.write": SharingWrite,
        "sharing.read": SharingRead,
        "file_requests.write": FileRequestsWrite,
        "file_requests.read": FileRequestsRead,
        "contacts.write": ContactsWrite,
        "contacts.read": ContactsRead,
        "openid": Openid,
        "profile": Profile,
        "email": Email,
    }

    PermissionIDs = map[Permission]int{
        AccountInfoWrite: 1,
        AccountInfoRead: 2,
        FilesMetadataWrite: 3,
        FilesMetadataRead: 4,
        FilesContentWrite: 5,
        FilesContentRead: 6,
        SharingWrite: 7,
        SharingRead: 8,
        FileRequestsWrite: 9,
        FileRequestsRead: 10,
        ContactsWrite: 11,
        ContactsRead: 12,
        Openid: 13,
        Profile: 14,
        Email: 15,
    }

    IdToPermission = map[int]Permission{
        1: AccountInfoWrite,
        2: AccountInfoRead,
        3: FilesMetadataWrite,
        4: FilesMetadataRead,
        5: FilesContentWrite,
        6: FilesContentRead,
        7: SharingWrite,
        8: SharingRead,
        9: FileRequestsWrite,
        10: FileRequestsRead,
        11: ContactsWrite,
        12: ContactsRead,
        13: Openid,
        14: Profile,
        15: Email,
    }
)

// ToString converts a Permission enum to its string representation
func (p Permission) ToString() (string, error) {
    if str, ok := PermissionStrings[p]; ok {
        return str, nil
    }
    return "", errors.New("invalid permission")
}

// ToID converts a Permission enum to its ID
func (p Permission) ToID() (int, error) {
    if id, ok := PermissionIDs[p]; ok {
        return id, nil
    }
    return 0, errors.New("invalid permission")
}

// PermissionFromString converts a string representation to its Permission enum
func PermissionFromString(s string) (Permission, error) {
    if p, ok := StringToPermission[s]; ok {
        return p, nil
    }
    return 0, errors.New("invalid permission string")
}

// PermissionFromID converts an ID to its Permission enum
func PermissionFromID(id int) (Permission, error) {
    if p, ok := IdToPermission[id]; ok {
        return p, nil
    }
    return 0, errors.New("invalid permission ID")
}
