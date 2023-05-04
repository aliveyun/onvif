

Enhanced and Improved from: [goonvif](https://github.com/yakovlevdmv/goonvif)
http://www.8fe.com/jiaocheng/7439.html
Onvif云台控制
云台控制说明：

x、y、z 范围都在0-1之间。
x为负数，表示左转，x为正数，表示右转。
y为负数，表示下转，y为正数，表示上转。
z为正数，表示拉近，z为负数，表示拉远。
通过x和y的组合，来实现云台的控制。
通过z的组合，来实现焦距控制。
func main() {
		//Getting an camera instance
		dev, err := onvif.NewDevice(onvif.DeviceParams{
			Xaddr:      "192.168.138.120:80",
			Username:   "admin",
			Password:   "abc12345",
			HttpClient: new(http.Client),
		})

 fmt.Println(dev ,err)
 //dev.PtzUp()
 dev.ControlPTZ(3,true,0.08)
 //time.Sleep(5 * time.Second)
}


void frmPtz::moveAbsolute()
{
    OnvifDevice *device = frm->getCurrentDevice();
    if (device) {
        QString profile = frm->getProfile();
        device->moveAbsolute(profile, x, y, z);
        frm->append(5, QString("执行绝对移动-> x: %1  y: %2  z: %3").arg(x).arg(y).arg(z));
    }
}

void frmPtz::moveRelative()
{
    OnvifDevice *device = frm->getCurrentDevice();
    if (device) {
        QString profile = frm->getProfile();
        device->moveRelative(profile, x, y, z);
        frm->append(5, QString("执行相对移动-> x: %1  y: %2  z: %3").arg(x).arg(y).arg(z));
    }
}

void frmPtz::setFrm(frmMain *frm)
{
    this->frm = frm;
}

void frmPtz::on_btnPtzUp_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.0;
        y = 0.1;
        z = 0.0;
        moveRelative();
    } else {
        y = 0.1;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzDown_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.0;
        y = -0.1;
        z = 0.0;
        moveRelative();
    } else {
        y = -0.0;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzLeft_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = -0.1;
        y = 0.0;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.0;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzRight_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.1;
        y = 0.0;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.1;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzLeftUp_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = -0.1;
        y = 0.1;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.0;
        y = 0.0;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzLeftDown_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = -0.1;
        y = -0.1;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.1;
        y = 0.1;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzRightUp_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.1;
        y = 0.1;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.0;
        y = 0.0;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzRightDown_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.1;
        y = -0.1;
        z = 0.0;
        moveRelative();
    } else {
        x = 0.1;
        y = 0.1;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzZoomIn_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.0;
        y = 0.0;
        z = 0.01;
        moveRelative();
    } else {
        z = 0.01;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzZoomOut_clicked()
{
    if (ui->rbtnMoveRelative->isChecked()) {
        x = 0.0;
        y = 0.0;
        z = -0.01;
        moveRelative();
    } else {
        z = 0.0;
        moveAbsolute();
    }
}

void frmPtz::on_btnPtzStop_clicked()
{
    OnvifDevice *device = frm->getCurrentDevice();
    if (device) {
        frm->setText("ptzStop");
        QString profile = frm->getProfile();
        device->moveStop(profile);
    }
}

void frmPtz::on_btnPtzReset_clicked()
{
    x = 0.0;
    y = 0.0;
    z = 0.0;
    moveAbsolute();
}
