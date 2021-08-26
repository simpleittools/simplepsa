<template>
  <form action="">
    <FirstName />
    <LastName />
    <CompanyName />
    <CellPhone />
    <DeskPhone />
    <Phone2 />
    <EmailAddress />
    <PhysicalAddress />
    <AddressAlt />
    <AddressState />
    <ZipCode />
    <InvoiceMethod />
    <button class="btn btn-success">Add</button>
  </form>
</template>

<script>
import AddressAlt from "@/components/formParts/AddressAlt";
import AddressState from "@/components/formParts/AddressState";
import CellPhone from "@/components/formParts/CellPhone";
import CompanyName from "@/components/formParts/CompanyName";
import DeskPhone from "@/components/formParts/DeskPhone";
import EmailAddress from "@/components/formParts/EmailAddress";
import FirstName from "@/components/formParts/FirstName";
import InvoiceMethod from "@/components/formParts/InvoiceMethod";
import LastName from "@/components/formParts/LastName";
import Phone2 from "@/components/formParts/Phone2";
import PhysicalAddress from "@/components/formParts/PhysicalAddress";
import ZipCode from "@/components/formParts/ZipCode";

export default {
  name: "WalkInAdd",
  components: {
    AddressAlt,
    AddressState,
    CellPhone,
    CompanyName,
    DeskPhone,
    EmailAddress,
    FirstName,
    InvoiceMethod,
    LastName,
    Phone2,
    PhysicalAddress,
    ZipCode,
  }
}

</script>

<style scoped>

</style>
