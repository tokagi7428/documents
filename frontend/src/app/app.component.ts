import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterOutlet } from '@angular/router';
import { AppService, DocumentItem } from './app.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet,FormsModule,CommonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  docs: DocumentItem[] = [];
  
  modalType: 'approve' | 'reject' | null = null;
  modalReason = '';

  constructor(private service: AppService) {
    this.service.getDocuments().subscribe((data) => {
      data.data.map(doc => this.docs.push(doc));
    });
  }

  getDocument(){
    // this.docs = [];
    this.service.getDocuments().subscribe((data) => {
      data.data.map(doc => this.docs.push(doc));
    });
  }

  get selectedIds() {
    return this.docs.filter(d => d.checked).map(d => d.id);
  }

  openApproveModal() {
    if (this.selectedIds.length === 0) {
      alert("กรุณาเลือกรายการที่ต้องการอนุมัติ");
      return;
    }
    this.modalReason = '';
    this.modalType = 'approve';
  }

  openRejectModal() {
    if (this.selectedIds.length === 0) {
      alert("กรุณาเลือกรายการที่ต้องการไม่อนุมัติ");
      return;
    }
    this.modalReason = '';
    this.modalType = 'reject';
  }

  confirm() {
    if (this.modalType === 'approve') {
      this.service.approve(this.selectedIds, this.modalReason).subscribe((data) => console.log(data));
    } else {
      this.service.reject(this.selectedIds, this.modalReason).subscribe((data) => console.log(data));
    }
    this.getDocument();
    this.closeModal();
  }

  closeModal() {
    this.modalType = null;
    this.modalReason = '';
    this.docs.forEach(d => d.checked = false);
  }

}
