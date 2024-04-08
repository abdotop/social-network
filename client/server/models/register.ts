export class Register {
    email: string;
    password: string;
    repeatPassword: string;
    firstName: string;
    lastName: string;
    dateOfBirth: Date;
    avatarImage?: string;
    nickname: string;
    aboutMe?: string;
    isPublic: boolean;
    [key: string]: any;

    constructor(data: any) {
        this.email = data.email;
        this.password = data.password;
        this.repeatPassword = data.repeatPassword;
        this.firstName = data.firstName;
        this.lastName = data.lastName;
        this.dateOfBirth = new Date(data.dateOfBirth);
        this.avatarImage = data.avatarImage !== undefined ? data.avatarImage : "uploads/default-avatar.png";
        this.nickname = data.nickname;
        this.aboutMe = data.aboutMe;
        this.isPublic = data.isPublic !== undefined ? data.isPublic : true;
    }

    static isEmail(email: string): boolean {
        const re = /\S+@\S+\.\S+/;
        return re.test(email);
    }

    validate(): [boolean, string] {
        const requiredFields = ['email', 'password', 'repeatPassword', 'firstName', 'lastName', 'dateOfBirth'];

        for (const field of requiredFields) {
            if (!this[field]) {
                return [false, `${field.charAt(0).toUpperCase() + field.slice(1)} is missing. Please provide it.`];
            }
        }

        if (!Register.isEmail(this.email)) {
            return [false, "Invalid email"];
        }

        if (this.password.length < 8) {
            return [false, "Password must be at least 8 characters long"];
        }

        if (this.password !== this.repeatPassword) {
            return [false, "Passwords do not match"];
        }

        return [true, "User data is valid"];
    }
}