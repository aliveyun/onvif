

Enhanced and Improved from: [goonvif](https://github.com/yakovlevdmv/goonvif)
http://www.8fe.com/jiaocheng/7439.html

3. 备注
需要特别注意：
1、import use-go/onvif 相关包的时候，一定要加上别名。
2、摄像头的地址一定要可以ping通，并且username和password不是登陆摄像头管理后台的用户，而是协议用户：

海康摄像头的配置路径：管理后台 -> 配置 -> 网络 -> 高级配置 -> 集成协议


 - interfacename: WLAN  # 网卡名称 或者"以太网" "eth0"等，使用ipconfig 或者 ifconfig 查看网卡名称   以太网 2


Onvif云台控制
云台控制说明：

x、y、z 范围都在0-1之间。
x为负数，表示左转，x为正数，表示右转。
y为负数，表示下转，y为正数，表示上转。
z为正数，表示拉近，z为负数，表示拉远。
通过x和y的组合，来实现云台的控制。
通过z的组合，来实现焦距控制。

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
