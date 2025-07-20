resource "aws_instance" "cd_target" {
  ami           = "ami-0150ccaf51ab55a51"         # Use an official or your own AMI ID
  instance_type = "t2.micro"
  subnet_id = "subnet-08cd3b75f943c88c4"

  tags = {
    Name = var.ec2_tag_name_value
  }
}
