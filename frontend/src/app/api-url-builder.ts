export class ApiUrlBuilder {
  private static readonly ORIGIN: string = document.location.origin;

  public static getApiUrlRequest(endpoint: string): string {
    return `${ApiUrlBuilder.ORIGIN}/api/${endpoint}`;
  }
}
