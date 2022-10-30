export class ResetModel {
    public token: string;
    public userId: string;
    public password: string;
    public confirmPassword: string;
}

export class ResetResultModel {
    public success: boolean;
    public message: string;
}
