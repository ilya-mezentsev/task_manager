export class ApiRequestBuilder {
  public static getLoginRequest(login: string, password: string): {user_name: string, user_password: string} {
    return {user_name: login, user_password: password};
  }
}
