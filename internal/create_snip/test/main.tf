variable "nyan" {
  type = string
}
variable "nyan1" {}
variable "nyan2" {}
variable "nyan3" {}
variable "nyan4" {
  default = "hoge"
}
variable "nyan5" {
  default = {
    tanuki = "tagarogu"
    iruka  = "statoko"
  }
  type = object({
    tanuki = string
    iruka  = string
  })
}
output "neko" {
  value = "neko"
}
