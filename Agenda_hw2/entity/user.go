package entity

type User struct {
	Name, Password, Email, Phone string
}
func (mUser User) init(t_Name, t_Password, t_Email, t_Phone string) {
	mUser.Name= t_Name
	mUser.Password= t_Password
	mUser.Email= t_Email
	mUser.Phone= t_Phone
}

/**
* @brief copy constructor
*/
func (mUser User) CopyUser(t_user User) {
	mUser.Name= t_user.Name
	mUser.Password= t_user.Password
	mUser.Email= t_user.Email
	mUser.Phone= t_user.Phone
}

/**
* @brief get the name of the user
* @return   return a string indicate the name of the user
*/
func (mUser User) GetName() string {
	return mUser.Name;
}

/**
* @brief set the name of the user
* @param   a string indicate the new name of the user
*/
func (mUser User) SetName(t_name string) {
	mUser.Name = t_name;
}

/**
* @brief get the password of the user
* @return   return a string indicate the password of the user
*/
func (mUser User) GetPassword() string {
	return mUser.Password;
}

/**
* @brief set the password of the user
* @param   a string indicate the new password of the user
*/
func (mUser User) SetPassword(t_password string) {
	mUser.Password = t_password;
}

/**
* @brief get the email of the user
* @return   return a string indicate the email of the user
*/
func (mUser User) GetEmail() string {
	return mUser.Email;
}

/**
* @brief set the email of the user
* @param   a string indicate the new email of the user
*/
func (mUser User) SetEmail(t_email string) {
	mUser.Email = t_email;
}

/**
* @brief get the phone of the user
* @return   return a string indicate the phone of the user
*/
func (mUser User) GetPhone() string {
	return mUser.Phone;
}

func (mUser User) SetPhone(t_phone string) {
	mUser.Phone = t_phone;
}