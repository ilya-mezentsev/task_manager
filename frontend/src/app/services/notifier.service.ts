import { Injectable } from '@angular/core';

declare namespace M {
  const toast: any;
}

@Injectable({
  providedIn: 'root'
})
export class NotifierService {

  constructor() {
  }

  public send(text: string): void {
    M.toast({html: text});
  }
}
