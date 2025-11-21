
type Role {
    id: String @1
    name: String @2
    description: String @3

    domain: String @4


    policies: [Policy] @10

    create_time: Timestamp @1000
    update_time: Timestamp @1001
}