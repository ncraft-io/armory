

type Subject {
    id: String @1
    name: String @2
    type: String @3 //< one of user,group,access_key
    domain: String @4

    attributes: Object @10 //< attributes of the subject for the policy, like: {department: "Finance"}
}
