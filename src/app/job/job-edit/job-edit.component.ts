import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from 'src/app/request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { isIncludeAdmin } from 'src/app/base/public';
@Component({
    selector: 'app-job-edit',
    templateUrl: './job-edit.component.html',
    styleUrls: ['./job-edit.component.scss'],
})
export class JobEditComponent implements OnInit {
    group!: FormGroup;
    id: any = 0;
    add: string = 's';
    @ViewChild('parametersChild') parametersChild: any;
    listOfOption: any[] = [
        { value: 'aa', label: 11 },
        { value: 'bb', label: 12 },
    ];
    productList: any[] = [
        { value: 'cc', label: 21 },
        { value: 'dd', label: 22 },
    ];
    url = '/app/history/api/';
    constructor(
        private fb: FormBuilder,
        private router: Router,
        private route: ActivatedRoute,
        private rs: RequestService,
        private msg: NzMessageService
    ) {}

    ngOnInit(): void {
        if (this.route.snapshot.paramMap.has('id')) {
            this.id = this.route.snapshot.paramMap.get('id');
            this.rs.get(this.url + `job/${this.id}`).subscribe((res) => { 
                //let data = res.data;
                this.build(res.data);
            });
        }

        this.build();

        this.rs
            .post('api/product/search', {})
            .subscribe((res) => {
                const data: any[] = [];

                res.data.filter((item: { id: string; name: string }) =>
                    data.push({
                        label: item.id + ' / ' + item.name,
                        value: item.id,
                    })
                );
                this.productList = data;
            })
            .add(() => {});
    }

    build(obj?: any) {
        obj = obj || {};
        this.group = this.fb.group({
            id: [obj.id || '', []],
            product_id: [obj.product_id || '', []],
            aggregators: [obj.aggregators || [], []],
            crontab: [obj.crontab || '', []],
            desc: [obj.desc || '', []],
            name: [obj.name || '', []],
            assign: [obj.assign || '', []],
            expression: [obj.expression || '', []],
            type: [obj.type || '', []],
        });
    }

    parameterslistData = [
        {
          title: '赋值',
          keyName: 'assign'
        }, {
          title: '表达式',
          keyName: 'expression'
        }, 
        {
          title: '聚合算法',
          keyName: 'type',
          type: 'select',
          listOfOption: 
          [{
            label: 'inc',
            value: 'inc'
          }, {
            label: 'dec',
            value: 'dec'
          },
          {
            label: 'avg',
            value: 'avg'
          },
          {
            label: 'count',
            value: 'count'
          },
          {
            label: 'min',
            value: 'min'
          },
          
          {
            label: 'max',
            value: 'max'
          },
          

          {
            label: 'sum',
            value: 'sum'
          }
          ,
          {
            label: 'last',
            value: 'last'
          }
          ,
          {
            label: 'first',
            value: 'first'
          }]
        } 
      ] 

      
      parameterAdd($event:any){

        $event.stopPropagation()
    if (this.parametersChild) {
      this.parametersChild.group.get('keyName').controls.unshift(
        this.fb.group({
            assign: ['', []],
            expression: ['', []],
         type: ['first', []] 
        })
      )
    }
     
      }
      getValueData() {
         
        const parametersGroup = this.parametersChild.group;
        const parameters = parametersGroup.get('keyName').controls.map((item: { value: any; }) => item.value);
        return    parameters   ;
      
      }
    submit() {
        let value = this.group.value;
        if (this.group.valid) {
            let url = this.id
                ? this.url + `job/${this.id}`
                : this.url + `job/create`;

           const sendData = JSON.parse(JSON.stringify(this.group.value));
           sendData.aggregators = this.getValueData();
        
            this.rs.post(url, sendData).subscribe((res) => {
                this.handleCancel();
                this.msg.success('保存成功');
            });

            return;
        } else {
            Object.values(this.group.controls).forEach((control) => {
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
