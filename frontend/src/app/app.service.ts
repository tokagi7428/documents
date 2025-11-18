import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

export interface DocumentItem {
  id: number;
  name: string;
  reason: string;
  status: 'รออนุมัติ' | 'อนุมัติ' | 'ไม่อนุมัติ';
  checked?: boolean;
}

@Injectable({ providedIn: 'root' })
export class AppService {
  private readonly url = `http://localhost:8080`

  constructor(private http: HttpClient) { }

  getDocuments() { return this.http.get<{data: DocumentItem[]}>(this.url+'/api/documents'); }

  approve(ids: number[], reason: string) {
    console.log({ids,reason});
    
    return this.http.post(this.url+'/api/documents/approve', { ids, reason });
  }

  reject(ids: number[], reason: string) {
    console.log({ids,reason});
    return this.http.post(this.url+'/api/documents/reject', { ids, reason });
  }

}
