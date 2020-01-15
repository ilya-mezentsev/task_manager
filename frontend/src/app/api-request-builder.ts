export class ApiRequestBuilder {
  public static getLoginRequest(login: string, password: string): {userName: string, userPassword: string} {
    return {userName: login, userPassword: password};
  }
}
