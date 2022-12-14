/**
 * @file-desc: ormConfig 可根据环境提供不同的数据库连接配置
 * @author: huanghe
 */
import { DataSource } from 'typeorm'
import { User } from './model/User'
import { Hospital } from './model/Hospital'
import { Department } from './model/Department'
import { AppointmentRecord } from './model/AppointmentRecord'
import { AppointmentSuccessRecord } from './model/AppointmentSuccessRecord'

export const OrmDataSource = new DataSource({
  type: 'mysql',
  host: 'localhost',
  port: 3306,
  username: 'root',
  password: '12345678',
  database: 'test',
  synchronize: true,
  logging: false,
  entities: [User, Department, Hospital, AppointmentRecord, AppointmentSuccessRecord]
})
