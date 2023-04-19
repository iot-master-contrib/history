 
  import { Component } from '@angular/core';
  import { Router } from '@angular/router';
  import { RequestService } from 'src/app/request.service'; 
  import { NzMessageService } from 'ng-zorro-antd/message';
  import { NzTableQueryParams } from 'ng-zorro-antd/table';
  import { NzModalService } from 'ng-zorro-antd/modal';
  import { ParseTableQuery } from 'src/app/base/table';  
  import {
      tableHeight,
      onAllChecked,
      onItemChecked,
      batchdel,
      refreshCheckedStatus,
  } from 'src/app/base/public';
  
  @Component({
    selector: 'app-jobs',
    templateUrl: './jobs.component.html',
    styleUrls: ['./jobs.component.scss']
  })
  export class JobsComponent {
      loading = true;
      datum: any[] = [];
      total = 1;
      pageSize = 20;
      uploading: Boolean = false;
      pageIndex = 1;
      query: any = {};
      url = '/app/history/api/';
      href!: string;
      filterLevel = [
          { text: '1', value: 1 },
          { text: '2', value: 2 },
          { text: '3', value: 3 },
      ];
      checked = false;
      indeterminate = false;
      setOfCheckedId = new Set<number>();
      delResData: any = [];
  
      constructor(
          private router: Router,
          private rs: RequestService,
          private modal: NzModalService,
          private msg: NzMessageService
      ) {
          //this.load();
      }
  
      reload() {
          this.datum = [];
          this.load();
      }
  
      load() {
          this.loading = true;
          this.rs
              .post(this.url + 'job/search', this.query)
              .subscribe((res) => {
                  const { data, total } = res;
                  this.datum = data || [];
                  this.total = total || 0;
                  this.setOfCheckedId.clear();
                  refreshCheckedStatus(this);
              })
              .add(() => {
                  this.loading = false;
              });
      }
  
      edit(id: any) {
          const path = `/job/edit/${id}`;
          this.router.navigateByUrl(path);
      }
      handleNew() {
          const path = `/job/create`;
          this.router.navigateByUrl(path);
      }
  
      delete(id: number, size?: number) {
          this.rs.get(this.url + `job/${id}/delete`).subscribe((res) => {
              if (!size) {
                  this.msg.success('删除成功');
                  this.datum = this.datum.filter((d) => d.id !== id);
              } else if (size) {
                  this.delResData.push(res);
                  if (size === this.delResData.length) {
                      this.msg.success('删除成功');
                      this.load();
                  }
              }
          });
      }
  
      onQuery($event: NzTableQueryParams) {
          ParseTableQuery($event, this.query);
          this.load();
      }
  
      pageIndexChange(pageIndex: number) {
          console.log('pageIndex:', pageIndex);
      }
  
      pageSizeChange(pageSize: number) {
          this.query.limit = pageSize;
          this.load();
      }
  
      search($event: string) { console.log()
          this.query.keyword = {
              name: $event,
           //   Message: $event,
          };
          
          this.query.skip = 0;
          this.load();
      }
  
      handleImport(e: any) {
          const file: File = e.target.files[0];
          const formData = new FormData();
          formData.append('file', file);
          this.rs
              .post(this.url + `job/import`, formData)
              .subscribe((res) => {
                  console.log(res);
              });
      }
  
      handleExport() {
          this.href = this.url + `job/export`;
      }
  
      read(data: any) {
          
      }
  
      disable(mess: number, id: any) {
          if (mess)
              this.rs
                  .get(this.url + `job/${id}/disable`)
                  .subscribe((res) => {
                      this.reload();
                  });
          else
              this.rs
                  .get(this.url + `job/${id}/enable`)
                  .subscribe((res) => {
                      this.reload();
                  });
      }
  
      getTableHeight() {
          return tableHeight(this);
      }
  
      handleBatchDel() {
          batchdel(this);
      }
  
      handleAllChecked(id: any) {
          onAllChecked(id, this);
      }
  
      handleItemChecked(id: number, checked: boolean) {
          onItemChecked(id, checked, this);
      }
  }
  