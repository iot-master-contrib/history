 
 
  import { Component, OnInit } from '@angular/core';
  import { FormBuilder, Validators, FormGroup } from "@angular/forms";
  import { ActivatedRoute, Router } from "@angular/router";
  import { RequestService } from "src/app/request.service";
  import { NzMessageService } from "ng-zorro-antd/message"; 
  import { isIncludeAdmin } from 'src/app/base/public';
  @Component({
    selector: 'app-job-edit',
    templateUrl: './job-edit.component.html',
    styleUrls: ['./job-edit.component.scss']
  })
  export class JobEditComponent  implements OnInit {
    group!: FormGroup;
    id: any = 0
    add:string='s'
    listOfOption:any[]=[{value:'aa',label:11},{value:'bb',label:12}]
    productList:any[]=[{value:'cc',label:21},{value:'dd',label:22}]
    url = '/app/history/api/';
    constructor(private fb: FormBuilder,
      private router: Router,
      private route: ActivatedRoute,
      private rs: RequestService,
      private msg: NzMessageService) {
    }
  
  
    ngOnInit(): void {
      if (this.route.snapshot.paramMap.has("id")) {
        this.id = this.route.snapshot.paramMap.get("id");
        this.rs.get(this.url+`job/${this.id}`).subscribe(res => {
          //let data = res.data;
          this.build(res.data)
        })
  
      }
  
      this.build()

      this.rs
      .post('api/product/search', {})
      .subscribe((res) => { 
        const data: any[] = [];

        res.data.filter((item: { id: string; name: string }) =>
          data.push({ label: item.id + ' / ' + item.name, value: item.id })
        );
        this.productList = data;
      })
      .add(() => {});
    }
  
    build(obj?: any) {
      obj = obj || {}
      this.group = this.fb.group({ 
         
        id: [obj.id || '', []],   
        product_id: [obj.product_id || '' ,[]], 
        aggregators: [obj.aggregators||[], []],  
        crontab: [obj.crontab || "" ,[]], 
        desc: [obj.desc || "" ,[]], 
        name: [obj.name || "", []],    
      })
    }
   
    submit() {
  
      if (this.group.valid) { 
        let url = this.id ? this.url+`job/${this.id}` : this.url+`job/create`;  
        if(this.group.value.aggregators==='')this.group.patchValue({aggregators:[]})
        this.group.patchValue({crontab:this.group.value.crontab+this.add})
        this.rs.post(url, this.group.value).subscribe(res => { 
          this.handleCancel()
          this.msg.success("保存成功")
        })
  
        return;
      }
      else {
        Object.values(this.group.controls).forEach(control => {
          if (control.invalid) {
            control.markAsDirty();
            control.updateValueAndValidity({ onlySelf: true });
          }
        });
  
      }
    }
    handleCancel() {
      const path = `/job`;
      this.router.navigateByUrl(path);
    }
  
  }
  
