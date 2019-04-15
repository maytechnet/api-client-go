# Go API client for quatrix

Download and upload files or folders, share them with predefined security options, manage your account and profile settings and a lot more functionalities can be easily integrated into your application using our Quatrix APIs. Learn more how to authenticate the Quatrix session, how to construct JSON formatted API calls and what responses to expect in our [API Guide](https://docs.maytech.net/display/MD/API+Guide).

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./quatrix"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.quatrix.it/api/1.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountApi* | [**AccountInfoGet**](docs/AccountApi.md#accountinfoget) | **Get** /account/info | Get account usage info
*AccountApi* | [**AccountListGet**](docs/AccountApi.md#accountlistget) | **Get** /account/list | List user accounts
*AccountApi* | [**AccountLogoGet**](docs/AccountApi.md#accountlogoget) | **Get** /account/logo | Get account logo
*AccountApi* | [**AccountMetadataGet**](docs/AccountApi.md#accountmetadataget) | **Get** /account/metadata | Get account metadata
*AccountApi* | [**AccountRolesGet**](docs/AccountApi.md#accountrolesget) | **Get** /account/roles | Call users of the account
*ActivityLogApi* | [**TrackingActivityGet**](docs/ActivityLogApi.md#trackingactivityget) | **Get** /tracking/activity | Get activity log
*ActivityLogApi* | [**TrackingCsvGet**](docs/ActivityLogApi.md#trackingcsvget) | **Get** /tracking/csv | Download CSV file with activity log
*ActivityLogApi* | [**TrackingDownloadsIdGet**](docs/ActivityLogApi.md#trackingdownloadsidget) | **Get** /tracking/downloads/{id} | Get shared file downloads
*ActivityLogApi* | [**TrackingFilesIdGet**](docs/ActivityLogApi.md#trackingfilesidget) | **Get** /tracking/files/{id} | Get share action files
*AuthApi* | [**SessionKeepaliveGet**](docs/AuthApi.md#sessionkeepaliveget) | **Get** /session/keepalive | Refresh session expiration time
*AuthApi* | [**SessionLoginGet**](docs/AuthApi.md#sessionloginget) | **Get** /session/login | Log in and get session ID details
*AuthApi* | [**SessionLoginPost**](docs/AuthApi.md#sessionloginpost) | **Post** /session/login | Log in to the account using MFA
*AuthApi* | [**SessionLogoutGet**](docs/AuthApi.md#sessionlogoutget) | **Get** /session/logout | Close the user&#39;s session
*AuthApi* | [**SessionUnblockCaptchaPost**](docs/AuthApi.md#sessionunblockcaptchapost) | **Post** /session/unblock-captcha | Unblock the session using CAPTCHA
*AutomationApi* | [**AutomationCreatePost**](docs/AutomationApi.md#automationcreatepost) | **Post** /automation/create | Create a new automation
*AutomationApi* | [**AutomationDeletePost**](docs/AutomationApi.md#automationdeletepost) | **Post** /automation/delete | Delete automations
*AutomationApi* | [**AutomationEditPost**](docs/AutomationApi.md#automationeditpost) | **Post** /automation/edit/ | Edit an existing automation
*AutomationApi* | [**AutomationGet**](docs/AutomationApi.md#automationget) | **Get** /automation | List all automations
*AutomationApi* | [**AutomationMetadataIdGet**](docs/AutomationApi.md#automationmetadataidget) | **Get** /automation/metadata/{id} | Get automation metadata
*BillingApi* | [**BillingUpgradePost**](docs/BillingApi.md#billingupgradepost) | **Post** /billing/upgrade | Upgrade the user&#39;s account
*ContactApi* | [**ContactCreatePost**](docs/ContactApi.md#contactcreatepost) | **Post** /contact/create | Create a new contact
*ContactApi* | [**ContactDeletePost**](docs/ContactApi.md#contactdeletepost) | **Post** /contact/delete | Delete a contact
*ContactApi* | [**ContactEditIdPost**](docs/ContactApi.md#contacteditidpost) | **Post** /contact/edit/{id} | Edit contact metadata
*ContactApi* | [**ContactGet**](docs/ContactApi.md#contactget) | **Get** /contact | List user contacts
*ContactApi* | [**ContactGroupGet**](docs/ContactApi.md#contactgroupget) | **Get** /contact/group | List available contact groups.
*ContactApi* | [**ContactMetadataIdGet**](docs/ContactApi.md#contactmetadataidget) | **Get** /contact/metadata/{id} | Get contact metadata
*ContactApi* | [**ContactPgpKeyIdGet**](docs/ContactApi.md#contactpgpkeyidget) | **Get** /contact/pgp-key/{id} | Get contact&#39;s PGP key
*ContactApi* | [**ContactUpgradeIdGet**](docs/ContactApi.md#contactupgradeidget) | **Get** /contact/upgrade/{id} | Upgrade a contact
*FileApi* | [**FileAddTagIdPost**](docs/FileApi.md#fileaddtagidpost) | **Post** /file/add-tag/{id} | Add a file tag
*FileApi* | [**FileCopyPost**](docs/FileApi.md#filecopypost) | **Post** /file/copy | Copy files
*FileApi* | [**FileCsvIdGet**](docs/FileApi.md#filecsvidget) | **Get** /file/csv/{id} | Download CSV file with Folder Content
*FileApi* | [**FileDeletePost**](docs/FileApi.md#filedeletepost) | **Post** /file/delete | Delete files
*FileApi* | [**FileDiffIdGet**](docs/FileApi.md#filediffidget) | **Get** /file/diff/{id} | Display changes of the file
*FileApi* | [**FileDownloadIdGet**](docs/FileApi.md#filedownloadidget) | **Get** /file/download/{id} | Download file
*FileApi* | [**FileDownloadLinkPost**](docs/FileApi.md#filedownloadlinkpost) | **Post** /file/download-link | Get download link
*FileApi* | [**FileInfoIdGet**](docs/FileApi.md#fileinfoidget) | **Get** /file/info/{id} | Get file info
*FileApi* | [**FileMakedirPost**](docs/FileApi.md#filemakedirpost) | **Post** /file/makedir | Create a new folder
*FileApi* | [**FileMetadataGet**](docs/FileApi.md#filemetadataget) | **Get** /file/metadata | Get metadata of files
*FileApi* | [**FileMetadataIdGet**](docs/FileApi.md#filemetadataidget) | **Get** /file/metadata/{id} | Get all file metadata
*FileApi* | [**FileMetadataPost**](docs/FileApi.md#filemetadatapost) | **Post** /file/metadata | Modify file metadata
*FileApi* | [**FileModifyPost**](docs/FileApi.md#filemodifypost) | **Post** /file/modify | Get file modification link
*FileApi* | [**FileMovePost**](docs/FileApi.md#filemovepost) | **Post** /file/move | Move files
*FileApi* | [**FilePreviewIdGet**](docs/FileApi.md#filepreviewidget) | **Get** /file/preview/{id} | Get a file preview
*FileApi* | [**FileRenameIdPost**](docs/FileApi.md#filerenameidpost) | **Post** /file/rename/{id} | Rename a file
*FileApi* | [**FileSearchPost**](docs/FileApi.md#filesearchpost) | **Post** /file/search | Search files
*FileApi* | [**FileSizeIdGet**](docs/FileApi.md#filesizeidget) | **Get** /file/size/{id} | Get file size
*FileApi* | [**FileTagsIdGet**](docs/FileApi.md#filetagsidget) | **Get** /file/tags/{id} | Get a list of file tags
*GroupApi* | [**ContactGroupGet**](docs/GroupApi.md#contactgroupget) | **Get** /contact/group | List available contact groups.
*GroupApi* | [**GroupGet**](docs/GroupApi.md#groupget) | **Get** /group | List available user groups
*GroupApi* | [**GroupMetadataIdGet**](docs/GroupApi.md#groupmetadataidget) | **Get** /group/metadata/{id} | Get group metadata
*GroupApi* | [**UserGroupGet**](docs/GroupApi.md#usergroupget) | **Get** /user/group | List all user groups
*InboxApi* | [**InboxGet**](docs/InboxApi.md#inboxget) | **Get** /inbox | Get contact&#39;s share tracking
*LanguageApi* | [**LanguagesDefaultGet**](docs/LanguageApi.md#languagesdefaultget) | **Get** /languages/default | Get default language
*LanguageApi* | [**LanguagesGet**](docs/LanguageApi.md#languagesget) | **Get** /languages | List available languages
*MFAApi* | [**ProfileRemoveMfaPost**](docs/MFAApi.md#profileremovemfapost) | **Post** /profile/remove-mfa | Disable MFA for the logged-in user
*MFAApi* | [**ProfileSetMfaPost**](docs/MFAApi.md#profilesetmfapost) | **Post** /profile/set-mfa | Enable MFA for the logged-in user
*MFAApi* | [**SessionLoginPost**](docs/MFAApi.md#sessionloginpost) | **Post** /session/login | Log in to the account using MFA
*MFAApi* | [**UserRemoveMfaPost**](docs/MFAApi.md#userremovemfapost) | **Post** /user/remove-mfa | Disable MFA for users
*MFAApi* | [**UserResetMfaPost**](docs/MFAApi.md#userresetmfapost) | **Post** /user/reset-mfa | Update existing MFA settings for users
*MFAApi* | [**UserSetMfaPost**](docs/MFAApi.md#usersetmfapost) | **Post** /user/set-mfa | Enable MFA for users
*PGPApi* | [**ContactPgpKeyIdGet**](docs/PGPApi.md#contactpgpkeyidget) | **Get** /contact/pgp-key/{id} | Get contact&#39;s PGP key
*PGPApi* | [**KeyRequestMetadataIdGet**](docs/PGPApi.md#keyrequestmetadataidget) | **Get** /key-request/metadata/{id} | Get PGP key request metadata
*PGPApi* | [**KeyRequestRespondIdPost**](docs/PGPApi.md#keyrequestrespondidpost) | **Post** /key-request/respond/{id} | Respond to PGP key request
*PGPApi* | [**PgpKeyCreatePost**](docs/PGPApi.md#pgpkeycreatepost) | **Post** /pgp-key/create | Create PGP key
*PGPApi* | [**PgpKeyDeleteIdGet**](docs/PGPApi.md#pgpkeydeleteidget) | **Get** /pgp-key/delete/{id} | Delete PGP key
*PGPApi* | [**PgpKeyEditIdPost**](docs/PGPApi.md#pgpkeyeditidpost) | **Post** /pgp-key/edit/{id} | Edit PGP key
*PGPApi* | [**PgpKeyMetadataIdGet**](docs/PGPApi.md#pgpkeymetadataidget) | **Get** /pgp-key/metadata/{id} | Get PGP key metadata
*PGPApi* | [**PgpKeyRecipientsPost**](docs/PGPApi.md#pgpkeyrecipientspost) | **Post** /pgp-key/recipients | Get recipients&#39; PGP keys
*PGPApi* | [**PgpKeyRequestIdsGet**](docs/PGPApi.md#pgpkeyrequestidsget) | **Get** /pgp-key/request/{ids[]} | Request PGP key
*PGPApi* | [**UserPgpKeyIdGet**](docs/PGPApi.md#userpgpkeyidget) | **Get** /user/pgp-key/{id} | Get user&#39;s PGP key
*PasswordResetApi* | [**ResetPasswordMetadataIdGet**](docs/PasswordResetApi.md#resetpasswordmetadataidget) | **Get** /reset-password/metadata/{id} | Get password reset request metadata
*PasswordResetApi* | [**ResetPasswordRequestPost**](docs/PasswordResetApi.md#resetpasswordrequestpost) | **Post** /reset-password/request | Request password reset
*PasswordResetApi* | [**ResetPasswordResetIdPost**](docs/PasswordResetApi.md#resetpasswordresetidpost) | **Post** /reset-password/reset/{id} | Reset password
*PreviewApi* | [**FilePreviewIdGet**](docs/PreviewApi.md#filepreviewidget) | **Get** /file/preview/{id} | Get a file preview
*PreviewApi* | [**PreviewIdGet**](docs/PreviewApi.md#previewidget) | **Get** /preview/{id} | Get binary preview data
*ProfileApi* | [**Profile2faGenerateGet**](docs/ProfileApi.md#profile2fagenerateget) | **Get** /profile/2fa/generate | Generate a new 2FA code
*ProfileApi* | [**ProfileGet**](docs/ProfileApi.md#profileget) | **Get** /profile | Get profile metadata
*ProfileApi* | [**ProfileInfoGet**](docs/ProfileApi.md#profileinfoget) | **Get** /profile/info | Retrieve additional profile info
*ProfileApi* | [**ProfileRemoveMfaPost**](docs/ProfileApi.md#profileremovemfapost) | **Post** /profile/remove-mfa | Disable MFA for the logged-in user
*ProfileApi* | [**ProfileSetMfaPost**](docs/ProfileApi.md#profilesetmfapost) | **Post** /profile/set-mfa | Enable MFA for the logged-in user
*ProfileApi* | [**ProfileSetPasswordPost**](docs/ProfileApi.md#profilesetpasswordpost) | **Post** /profile/set-password | Change profile password
*ProfileApi* | [**ProfileSetPost**](docs/ProfileApi.md#profilesetpost) | **Post** /profile/set | Update profile metadata
*ProjectFolderApi* | [**ProjectFolderAddUsersIdPost**](docs/ProjectFolderApi.md#projectfolderaddusersidpost) | **Post** /project-folder/add-users/{id} | Add users to the project folder
*ProjectFolderApi* | [**ProjectFolderCreatePost**](docs/ProjectFolderApi.md#projectfoldercreatepost) | **Post** /project-folder/create | Create a project folder
*ProjectFolderApi* | [**ProjectFolderDeleteIdGet**](docs/ProjectFolderApi.md#projectfolderdeleteidget) | **Get** /project-folder/delete/{id} | Convert a project folder to a folder
*ProjectFolderApi* | [**ProjectFolderDeleteUsersPost**](docs/ProjectFolderApi.md#projectfolderdeleteuserspost) | **Post** /project-folder/delete-users/ | Remove project folder users
*ProjectFolderApi* | [**ProjectFolderEditUsersIdPost**](docs/ProjectFolderApi.md#projectfoldereditusersidpost) | **Post** /project-folder/edit-users/{id} | Update users’ permissions of the project folder
*ProjectFolderApi* | [**ProjectFolderGet**](docs/ProjectFolderApi.md#projectfolderget) | **Get** /project-folder | List available project folders for a logged-in user
*ProjectFolderApi* | [**ProjectFolderMetadataIdGet**](docs/ProjectFolderApi.md#projectfoldermetadataidget) | **Get** /project-folder/metadata/{id} | Get project folder metadata
*ProjectFolderApi* | [**ProjectFolderProjectUsersPost**](docs/ProjectFolderApi.md#projectfolderprojectuserspost) | **Post** /project-folder/project-users | List all project folders for given users
*ProjectFolderApi* | [**ProjectFolderSetUsersPost**](docs/ProjectFolderApi.md#projectfoldersetuserspost) | **Post** /project-folder/set-users | Add users to project folders.
*ProjectFolderApi* | [**ProjectFolderUsersIdGet**](docs/ProjectFolderApi.md#projectfolderusersidget) | **Get** /project-folder/users/{id} | List users of the project folder
*SSHKeyApi* | [**SshKeyCreatePost**](docs/SSHKeyApi.md#sshkeycreatepost) | **Post** /ssh-key/create | Create a new SSH key
*SSHKeyApi* | [**SshKeyDeletePost**](docs/SSHKeyApi.md#sshkeydeletepost) | **Post** /ssh-key/delete | Delete SSH key
*SSHKeyApi* | [**SshKeyEditPost**](docs/SSHKeyApi.md#sshkeyeditpost) | **Post** /ssh-key/edit | Edit SSH key metadata
*SSHKeyApi* | [**SshKeyGet**](docs/SSHKeyApi.md#sshkeyget) | **Get** /ssh-key | List available SSH keys
*SSHKeyApi* | [**SshKeyMetadataIdGet**](docs/SSHKeyApi.md#sshkeymetadataidget) | **Get** /ssh-key/metadata/{id} | Get SSH key metadata
*ServiceApi* | [**ServiceGet**](docs/ServiceApi.md#serviceget) | **Get** /service | List available services
*ServiceApi* | [**ServiceMetadataIdGet**](docs/ServiceApi.md#servicemetadataidget) | **Get** /service/metadata/{id} | Get service metadata
*ShareApi* | [**FilesReturnMakedirIdPost**](docs/ShareApi.md#filesreturnmakediridpost) | **Post** /files-return/makedir/{id} | Create a directory for returned files
*ShareApi* | [**FilesReturnMetadataIdGet**](docs/ShareApi.md#filesreturnmetadataidget) | **Get** /files-return/metadata/{id} | Get return files metadata
*ShareApi* | [**FilesReturnSendPost**](docs/ShareApi.md#filesreturnsendpost) | **Post** /files-return/send | Return files in the created share
*ShareApi* | [**FilesReturnUploadLinkIdPost**](docs/ShareApi.md#filesreturnuploadlinkidpost) | **Post** /files-return/upload-link/{id} | Get return files upload link
*ShareApi* | [**QuicklinkCreatePost**](docs/ShareApi.md#quicklinkcreatepost) | **Post** /quicklink/create | Create a quicklink
*ShareApi* | [**QuicklinkLoginPinPost**](docs/ShareApi.md#quicklinkloginpinpost) | **Post** /quicklink/login-pin | Log in with PIN to access a quicklink
*ShareApi* | [**QuicklinkRevokeIdGet**](docs/ShareApi.md#quicklinkrevokeidget) | **Get** /quicklink/revoke/{id} | Revoke a quicklink
*ShareApi* | [**ShareCreatePost**](docs/ShareApi.md#sharecreatepost) | **Post** /share/create | Create a file share
*ShareApi* | [**ShareDownloadIdGet**](docs/ShareApi.md#sharedownloadidget) | **Get** /share/download/{id} | Download share files
*ShareApi* | [**ShareDownloadInfoIdGet**](docs/ShareApi.md#sharedownloadinfoidget) | **Get** /share/download-info/{id} | Get share download info
*ShareApi* | [**ShareDownloadLinkIdGet**](docs/ShareApi.md#sharedownloadlinkidget) | **Get** /share/download-link/{id} | Get download link for all files
*ShareApi* | [**ShareDownloadLinkIdPost**](docs/ShareApi.md#sharedownloadlinkidpost) | **Post** /share/download-link/{id} | Get download link for specified files
*ShareApi* | [**ShareFilesIdGet**](docs/ShareApi.md#sharefilesidget) | **Get** /share/files/{id} | List shared files
*ShareApi* | [**ShareLoginPinPost**](docs/ShareApi.md#shareloginpinpost) | **Post** /share/login-pin | Log in with PIN to access a share
*ShareApi* | [**SharePreviewIdGet**](docs/ShareApi.md#sharepreviewidget) | **Get** /share/preview/{id} | Preview a shared file
*ShareApi* | [**ShareRecipientsGet**](docs/ShareApi.md#sharerecipientsget) | **Get** /share/recipients | List all contacts for the share
*ShareApi* | [**ShareRequestPost**](docs/ShareApi.md#sharerequestpost) | **Post** /share/request | Send a request to share files
*ShareApi* | [**ShareRevokeIdGet**](docs/ShareApi.md#sharerevokeidget) | **Get** /share/revoke/{id} | Revoke a share
*ShareApi* | [**ShareSendRequestIdPost**](docs/ShareApi.md#sharesendrequestidpost) | **Post** /share/send-request/{id} | Request files. Use /share/request API call instead.
*ShareApi* | [**TrackingGet**](docs/ShareApi.md#trackingget) | **Get** /tracking/ | List share actions metadata for all users
*ShareApi* | [**TrackingIdGet**](docs/ShareApi.md#trackingidget) | **Get** /tracking/{id} | List share actions metadata for a user
*SiteSettingsApi* | [**SettingsAuthMethodsGet**](docs/SiteSettingsApi.md#settingsauthmethodsget) | **Get** /settings/auth-methods | Get available authentication methods
*SiteSettingsApi* | [**SettingsGet**](docs/SiteSettingsApi.md#settingsget) | **Get** /settings | Get site settings
*SiteSettingsApi* | [**SettingsSetPost**](docs/SiteSettingsApi.md#settingssetpost) | **Post** /settings/set | Set site settings
*SiteSettingsApi* | [**SettingsUploadLogoLinkGet**](docs/SiteSettingsApi.md#settingsuploadlogolinkget) | **Get** /settings/upload-logo-link | Get a new logo upload link
*UploadApi* | [**FileModifyPost**](docs/UploadApi.md#filemodifypost) | **Post** /file/modify | Get file modification link
*UploadApi* | [**SettingsUploadLogoLinkGet**](docs/UploadApi.md#settingsuploadlogolinkget) | **Get** /settings/upload-logo-link | Get a new logo upload link
*UploadApi* | [**UploadFinalizeIdGet**](docs/UploadApi.md#uploadfinalizeidget) | **Get** /upload/finalize/{id} | Finalize chunked file upload
*UploadApi* | [**UploadLinkPost**](docs/UploadApi.md#uploadlinkpost) | **Post** /upload/link | Get file upload link
*UserApi* | [**UserCreatePost**](docs/UserApi.md#usercreatepost) | **Post** /user/create | Create a user
*UserApi* | [**UserDeletePost**](docs/UserApi.md#userdeletepost) | **Post** /user/delete | Delete users
*UserApi* | [**UserEditPost**](docs/UserApi.md#usereditpost) | **Post** /user/edit | Update metadata of users
*UserApi* | [**UserGet**](docs/UserApi.md#userget) | **Get** /user | List users
*UserApi* | [**UserGroupGet**](docs/UserApi.md#usergroupget) | **Get** /user/group | List all user groups
*UserApi* | [**UserMetadataIdGet**](docs/UserApi.md#usermetadataidget) | **Get** /user/metadata/{id} | Get user metadata
*UserApi* | [**UserPgpKeyIdGet**](docs/UserApi.md#userpgpkeyidget) | **Get** /user/pgp-key/{id} | Get user&#39;s PGP key
*UserApi* | [**UserRemoveMfaPost**](docs/UserApi.md#userremovemfapost) | **Post** /user/remove-mfa | Disable MFA for users
*UserApi* | [**UserResetMfaPost**](docs/UserApi.md#userresetmfapost) | **Post** /user/reset-mfa | Update existing MFA settings for users
*UserApi* | [**UserSetMfaPost**](docs/UserApi.md#usersetmfapost) | **Post** /user/set-mfa | Enable MFA for users
*UserApi* | [**UserSignupPost**](docs/UserApi.md#usersignuppost) | **Post** /user/signup | Register a new user
*WidgetApi* | [**WidgetFinalizeUploadIdGet**](docs/WidgetApi.md#widgetfinalizeuploadidget) | **Get** /widget/finalize-upload/{id} | Finalize chunked upload of the widget
*WidgetApi* | [**WidgetMetadataIdGet**](docs/WidgetApi.md#widgetmetadataidget) | **Get** /widget/metadata/{id} | Get all widget metadata
*WidgetApi* | [**WidgetUploadLinkIdPost**](docs/WidgetApi.md#widgetuploadlinkidpost) | **Post** /widget/upload-link/{id} | Get widget upload link


## Documentation For Models

 - [AccountInfoResp](docs/AccountInfoResp.md)
 - [AccountListRespItems](docs/AccountListRespItems.md)
 - [AccountMetadataResp](docs/AccountMetadataResp.md)
 - [AccountRolesRespItems](docs/AccountRolesRespItems.md)
 - [AutomationCreateReq](docs/AutomationCreateReq.md)
 - [AutomationDeleteResp](docs/AutomationDeleteResp.md)
 - [AutomationEditReq](docs/AutomationEditReq.md)
 - [AutomationOptions](docs/AutomationOptions.md)
 - [AutomationResp](docs/AutomationResp.md)
 - [BillingUpgradeReq](docs/BillingUpgradeReq.md)
 - [BillingUpgradeResp](docs/BillingUpgradeResp.md)
 - [ContactCreateReq](docs/ContactCreateReq.md)
 - [ContactDeleteRespItems](docs/ContactDeleteRespItems.md)
 - [ContactEditResp](docs/ContactEditResp.md)
 - [ContactGroupRespItems](docs/ContactGroupRespItems.md)
 - [ContactResp](docs/ContactResp.md)
 - [CopyMoveFilesReq](docs/CopyMoveFilesReq.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [FileAddTagReq](docs/FileAddTagReq.md)
 - [FileCsvResp](docs/FileCsvResp.md)
 - [FileDiffResp](docs/FileDiffResp.md)
 - [FileDiffRespData](docs/FileDiffRespData.md)
 - [FileInfoResp](docs/FileInfoResp.md)
 - [FileMetadataGetResp](docs/FileMetadataGetResp.md)
 - [FileMetadataPostReq](docs/FileMetadataPostReq.md)
 - [FileMetadataPostResp](docs/FileMetadataPostResp.md)
 - [FileModifyReq](docs/FileModifyReq.md)
 - [FileModifyResp](docs/FileModifyResp.md)
 - [FilePreviewResp](docs/FilePreviewResp.md)
 - [FileRenameReq](docs/FileRenameReq.md)
 - [FileRenameResp](docs/FileRenameResp.md)
 - [FileResp](docs/FileResp.md)
 - [FileSizeResp](docs/FileSizeResp.md)
 - [FileTagResp](docs/FileTagResp.md)
 - [FilesReturnMakedirReq](docs/FilesReturnMakedirReq.md)
 - [FilesReturnMakedirResps](docs/FilesReturnMakedirResps.md)
 - [FilesReturnMetadataResp](docs/FilesReturnMetadataResp.md)
 - [FilesReturnSendReq](docs/FilesReturnSendReq.md)
 - [FilesReturnSendResp](docs/FilesReturnSendResp.md)
 - [FilesReturnUploadLinkReq](docs/FilesReturnUploadLinkReq.md)
 - [FilesReturnUploadLinkResp](docs/FilesReturnUploadLinkResp.md)
 - [Group](docs/Group.md)
 - [GroupMetadata](docs/GroupMetadata.md)
 - [GroupResp](docs/GroupResp.md)
 - [IdName](docs/IdName.md)
 - [IdResp](docs/IdResp.md)
 - [IdsReq](docs/IdsReq.md)
 - [IdsResp](docs/IdsResp.md)
 - [InboxRespItems](docs/InboxRespItems.md)
 - [JobResp](docs/JobResp.md)
 - [KeyRequestMetadataResp](docs/KeyRequestMetadataResp.md)
 - [KeyRequestRespondReq](docs/KeyRequestRespondReq.md)
 - [LanguagesDefaultResp](docs/LanguagesDefaultResp.md)
 - [LanguagesRespItems](docs/LanguagesRespItems.md)
 - [MakeDirReq](docs/MakeDirReq.md)
 - [PfCreateResp](docs/PfCreateResp.md)
 - [PfDeleteUsersRespItems](docs/PfDeleteUsersRespItems.md)
 - [PfMetadataResp](docs/PfMetadataResp.md)
 - [PfSetUsersReq](docs/PfSetUsersReq.md)
 - [PfSetUsersRespItems](docs/PfSetUsersRespItems.md)
 - [PfUsersListReq](docs/PfUsersListReq.md)
 - [PfUsersListRespItems](docs/PfUsersListRespItems.md)
 - [PfaddUsersReq](docs/PfaddUsersReq.md)
 - [PfcreateReq](docs/PfcreateReq.md)
 - [PfdeleteUsersReq](docs/PfdeleteUsersReq.md)
 - [PfeditUsersReq](docs/PfeditUsersReq.md)
 - [PgpCreateReq](docs/PgpCreateReq.md)
 - [PgpEditReq](docs/PgpEditReq.md)
 - [PgpKeyRecipientsRespItems](docs/PgpKeyRecipientsRespItems.md)
 - [PgpKeyResp](docs/PgpKeyResp.md)
 - [ProfileInfoResp](docs/ProfileInfoResp.md)
 - [ProfileRemoveMfaReq](docs/ProfileRemoveMfaReq.md)
 - [ProfileRemoveMfaResp](docs/ProfileRemoveMfaResp.md)
 - [ProfileResp](docs/ProfileResp.md)
 - [ProfileRespShareTypes](docs/ProfileRespShareTypes.md)
 - [ProfileSetMfaReq](docs/ProfileSetMfaReq.md)
 - [ProfileSetMfaResp](docs/ProfileSetMfaResp.md)
 - [ProfileSetPasswordReq](docs/ProfileSetPasswordReq.md)
 - [ProfileSetPasswordResp](docs/ProfileSetPasswordResp.md)
 - [ProfileSetReq](docs/ProfileSetReq.md)
 - [ProfileSetResp](docs/ProfileSetResp.md)
 - [ProjectFoldersListRespItem](docs/ProjectFoldersListRespItem.md)
 - [ProjectOwner](docs/ProjectOwner.md)
 - [ProjectfolderdeleteusersUsers](docs/ProjectfolderdeleteusersUsers.md)
 - [ProjectfoldersetusersParams](docs/ProjectfoldersetusersParams.md)
 - [ProjectfoldersetusersUsers](docs/ProjectfoldersetusersUsers.md)
 - [QuicklinkCreateReq](docs/QuicklinkCreateReq.md)
 - [QuicklinkCreateResp](docs/QuicklinkCreateResp.md)
 - [QuicklinkLoginPinReq](docs/QuicklinkLoginPinReq.md)
 - [ResetPasswordMetadataResp](docs/ResetPasswordMetadataResp.md)
 - [ResetPasswordRequestReq](docs/ResetPasswordRequestReq.md)
 - [ResetPasswordRequestResp](docs/ResetPasswordRequestResp.md)
 - [ResetPasswordResetReq](docs/ResetPasswordResetReq.md)
 - [SearchQuery](docs/SearchQuery.md)
 - [SearchReq](docs/SearchReq.md)
 - [SearchTag](docs/SearchTag.md)
 - [ServiceResp](docs/ServiceResp.md)
 - [ServiceRespUsers](docs/ServiceRespUsers.md)
 - [SessionLoginPostResp](docs/SessionLoginPostResp.md)
 - [SessionLoginResp](docs/SessionLoginResp.md)
 - [SessionUnblockCaptchaResp](docs/SessionUnblockCaptchaResp.md)
 - [SettingsAuthMethodsRespItems](docs/SettingsAuthMethodsRespItems.md)
 - [SettingsResp](docs/SettingsResp.md)
 - [SettingsSetReq](docs/SettingsSetReq.md)
 - [SettingsUploadLogoLinkResp](docs/SettingsUploadLogoLinkResp.md)
 - [SettingssetShareTypes](docs/SettingssetShareTypes.md)
 - [ShareCreateReq](docs/ShareCreateReq.md)
 - [ShareCreateResp](docs/ShareCreateResp.md)
 - [ShareDownloadInfoResp](docs/ShareDownloadInfoResp.md)
 - [ShareDownloadLinkReq](docs/ShareDownloadLinkReq.md)
 - [ShareFilesRespItems](docs/ShareFilesRespItems.md)
 - [ShareLoginPinReq](docs/ShareLoginPinReq.md)
 - [ShareRecipientsResp](docs/ShareRecipientsResp.md)
 - [ShareRequestReq](docs/ShareRequestReq.md)
 - [ShareRequestResp](docs/ShareRequestResp.md)
 - [ShareSendRequestReq](docs/ShareSendRequestReq.md)
 - [ShortUserService](docs/ShortUserService.md)
 - [SshKeyCreateReq](docs/SshKeyCreateReq.md)
 - [SshKeyDeleteReq](docs/SshKeyDeleteReq.md)
 - [SshKeyEditReq](docs/SshKeyEditReq.md)
 - [SshKeyResp](docs/SshKeyResp.md)
 - [StringSearchQueryElement](docs/StringSearchQueryElement.md)
 - [TagSearchQueryElement](docs/TagSearchQueryElement.md)
 - [TimestampSearchQueryElement](docs/TimestampSearchQueryElement.md)
 - [TrackingActivityRespItems](docs/TrackingActivityRespItems.md)
 - [TrackingCsvRespItems](docs/TrackingCsvRespItems.md)
 - [TrackingDownloadsRespItems](docs/TrackingDownloadsRespItems.md)
 - [TrackingFilesRespItems](docs/TrackingFilesRespItems.md)
 - [TrackingIdRespItems](docs/TrackingIdRespItems.md)
 - [TrackingRespItems](docs/TrackingRespItems.md)
 - [UnblockCaptchaReq](docs/UnblockCaptchaReq.md)
 - [UploadFinalizeResp](docs/UploadFinalizeResp.md)
 - [UploadLinkReq](docs/UploadLinkReq.md)
 - [UserCreateReq](docs/UserCreateReq.md)
 - [UserDeleteReq](docs/UserDeleteReq.md)
 - [UserEditReq](docs/UserEditReq.md)
 - [UserPermissionReq](docs/UserPermissionReq.md)
 - [UserPermissionResp](docs/UserPermissionResp.md)
 - [UserRemoveMfaReq](docs/UserRemoveMfaReq.md)
 - [UserResetMfaReq](docs/UserResetMfaReq.md)
 - [UserResp](docs/UserResp.md)
 - [UserSetMfaReq](docs/UserSetMfaReq.md)
 - [UserSignupReq](docs/UserSignupReq.md)
 - [WidgetFinalizeUploadResp](docs/WidgetFinalizeUploadResp.md)
 - [WidgetUploadLinkReq](docs/WidgetUploadLinkReq.md)
 - [WidgetUploadLinkResp](docs/WidgetUploadLinkResp.md)


## Documentation For Authorization

## api_key
- **Type**: API key 

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: "APIKEY",
		Prefix: "Bearer", // Omit if not necessary.
	})
    r, err := client.Service.Operation(auth, args)
```
## basicAuth
- **Type**: HTTP basic authentication

Example
```
	auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
		UserName: "username",
		Password: "password",
	})
    r, err := client.Service.Operation(auth, args)
```

## Author



