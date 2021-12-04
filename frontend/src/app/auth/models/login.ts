export class LoginModel {
  public emailaddress: string;
  public password: string;
}

export class LoginResultModel {
  public success: boolean;
  public message: string;
  public uid: string;
  public sid: string;
}
