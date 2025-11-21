
type Permission {
    id: String @1
    name: String @2
    description: String @3

    domain: String @4

    resource: String @10
    action: String @11
    attributes: Object @12

    create_time: Timestamp @1000
    update_time: Timestamp @1001
}