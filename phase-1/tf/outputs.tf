output "ec2_public_ip" {
    value = aws_instance.ec2.public_ip
}

output "nacl_id" {
    value = aws_network_acl.nacl.id
}

output "sg_id" {
    value = aws_security_group.sg.id
}