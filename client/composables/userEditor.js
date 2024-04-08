export const changeAvatar = async (file) => {
  const store = useAuth();
  const userInfos = useAuthUser();
  const body = new FormData();
  body.append("file", file);
  
  return new Promise(async (resolve, reject) => {
    try {
      const response = await $fetch("/api/settings/updateavatar", {
        method: "PUT",
        body: body,
      });

      if (response.ok === false) {
        return response.message;
      }
      
      resolve(response);
    } catch (error) {
      reject(error);
    }
  });
};

export const editUser = async (user) => {
  const store = useAuth();
  const currentUserInfos = useAuthUser()
  
  const error = validateUserInfo(user);
  if (error) {
    return error;
  }


  return new Promise(async (resolve, reject) => {
    const data = {
      email: user.email.trim(),
      password: "",
      firstName: user.firstName.trim(),
      lastName: user.lastName.trim(),
      dateOfBirth: new Date(user.dateOfBirth),
      avatarImage: currentUserInfos.avatarImage,
      nickname: user.nickname.trim(),
      aboutMe: user.aboutMe.trim(),
      isPublic: user.isPublic === "public" ? true : false,
    };

    try {
      const response = await $fetch("/api/updateuser", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      store.setUser({ ...response.user, isLoggedIn: true });
      resolve(response.message);
    } catch (error) {
      reject(error);
    }
  });
};

export const updatePassword = async (password) => {
  const store = useAuth();
  const user = useAuthUser();

  const error = validateUpdatePassword(password.currentPassword, password.newPassword, password.repeatNewPassword)
  if (error) {
    return error
  }

  return new Promise(async (resolve, reject) => {
    const data = {
      email: user.email,
      password: password.currentPassword.trim(),
      newPassword: password.newPassword.trim(),
    };

    try {
      const response = await $fetch("/api/updatepassword", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      store.setUser({ ...response.user, isLoggedIn: true });
      resolve(response.message);
    } catch (error) {
      reject(error);
    }
  });
};

function validateUserInfo(userInfo) {

  // Validate nickname
  // if (!userInfo.nickname || userInfo.nickname.trim() === "") {
  //   return "Nickname is required";
  // }

  // Validate firstName
  if (!userInfo.firstName || userInfo.firstName.trim() === "") {
    return "First name is required";
  }

  // Validate lastName
  if (!userInfo.lastName || userInfo.lastName.trim() === "") {
    return "Last name is required";
  }

  // Validate email
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!userInfo.email || !emailRegex.test(userInfo.email)) {
    return "Invalid email address";
  }

  // Validate dateOfBirth
  // const currentDate = new Date();
  // const minDateOfBirth = new Date(currentDate.getFullYear() - 10, currentDate.getMonth(), currentDate.getDate());
  // const dateOfBirth = new Date(userInfo.dateOfBirth);
  // if (!userInfo.dateOfBirth || dateOfBirth > minDateOfBirth) {
  //     return 'You must be at least 10 years old';
  // }

  if (typeof userInfo.aboutMe != "string") {
    return "About You must be text";
  }

  // Validate state
  if (userInfo.isPublic !== 'public' && userInfo.isPublic !== 'private') {
    return "Profile status must be public or private";
  }

  return null;
}

function validateUpdatePassword(
  currentPassword,
  newPassword,
  repeatNewPassword
) {

  // Check if current password is provided
  if (!currentPassword.trim() || !newPassword.trim() || !repeatNewPassword.trim()) {
    return "All fields are required";
  }
  
    // Check if new password matches the repeat new password
  if (newPassword !== repeatNewPassword) {
    return "New password and repeat new password do not match";
  }

  // Check if new password is the same as the current password
  if (newPassword === currentPassword) {
    return "New password must be different from the current password";
  }
  
  // Password length should be at least 8 characters
  if (newPassword.length < 8) {
    return "New password must be at least 8 characters long";
  }
  
  // Password should contain at least one lowercase letter, one uppercase letter, one digit, and one special character
  // const regex =
  //   /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/;
  // if (!regex.test(newPassword)) {
  //   return "New password must contain at least one lowercase letter, one uppercase letter, one digit, and one special character";
  // }
  // Password is valid
  return null;
}
