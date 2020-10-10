<template>
    <div>
        <b-navbar toggleable="lg" type="dark" variant="info">
            <b-button class="button-burger" v-b-toggle.sidebar-left>
                <b-icon icon="justify" variant="white"></b-icon>
            </b-button>
            <b-sidebar id="sidebar-left" bg-variant="dark" text-variant="light" title="Genos" left>
                <div class="px-3 py-2">
                    <SidebarContent></SidebarContent>
                </div>
            </b-sidebar>

            <b-navbar-toggle class="button-burger" target="nav-collapse">
                <template v-slot:default="{ expanded }">
                    <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon>
                    <b-icon v-else icon="chevron-bar-down"></b-icon>
                </template>
            </b-navbar-toggle>

            <b-collapse id="nav-collapse" is-nav>
                <b-navbar-nav class="ml-auto">
                    <b-nav-form>
                        <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
                        <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
                    </b-nav-form>
                </b-navbar-nav>
            </b-collapse>
            <b-button v-b-tooltip.hover title="Settings" class="button-burger">
                <b-icon icon="gear" variant="white"></b-icon>
            </b-button>
            <b-button v-b-tooltip.hover title="Profile" class="button-burger">
                <b-icon icon="person" variant="white"></b-icon>
            </b-button>
            <b-button v-b-tooltip.hover title="Logout" class="button-burger">
                <b-icon icon="chevron-bar-right" variant="white"></b-icon>
            </b-button>
        </b-navbar>
    </div>
</template>

<script>
    import SidebarContent from "./SidebarContent";
    export default {
        components: { SidebarContent },
        name: "HeaderTest"
    }
</script>

<style scoped lang="scss">
    .button-burger{
        background-color: #17a2b8 ;
        border: none;
        border-radius: 0;
    }
    .button-burger:hover{
        background-color: #17a2b8 ;
        svg{
            color: #fc2f05 !important;
        }
    }
    .navbar-toggler {
        border: 0;
        outline: none !important;
        box-shadow: none;
        svg{
            color: white;
        }
    }
</style>