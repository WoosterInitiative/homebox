<script setup lang="ts">
  import { Detail } from "~~/components/global/DetailsSection/types";
  import { themes } from "~~/lib/data/themes";
  import { CurrenciesCurrency, NotifierCreate, NotifierOut } from "~~/lib/api/types/data-contracts";

  definePageMeta({
    middleware: ["auth"],
  });
  useHead({
    title: "Homebox | Profile",
  });

  const api = useUserApi();
  const confirm = useConfirm();
  const notify = useNotifier();

  const currencies = computedAsync(async () => {
    const resp = await api.group.currencies();
    if (resp.error) {
      notify.error("Failed to get currencies");
      return [];
    }

    return resp.data;
  });

  // Currency Selection
  const currency = ref<CurrenciesCurrency>({
    code: "USD",
    name: "United States Dollar",
    local: "en-US",
    symbol: "$",
  });
  watch(currency, () => {
    if (group.value) {
      group.value.currency = currency.value.code;
    }
  });

  const currencyExample = computed(() => {
    const formatter = new Intl.NumberFormat("en-US", {
      style: "currency",
      currency: currency.value ? currency.value.code : "USD",
    });

    return formatter.format(1000);
  });

  const { data: group } = useAsyncData(async () => {
    const { data } = await api.group.get();
    return data;
  });

  // Sync Initial Currency
  watch(group, () => {
    if (!group.value) {
      return;
    }

    // @ts-expect-error - typescript is stupid, it should know group.value is not null
    const found = currencies.value.find(c => c.code === group.value.currency);
    if (found) {
      currency.value = found;
    }
  });

  async function updateGroup() {
    if (!group.value) {
      return;
    }

    const { data, error } = await api.group.update({
      name: group.value.name,
      currency: group.value.currency,
    });

    if (error) {
      notify.error("Failed to update group");
      return;
    }

    group.value = data;
    notify.success("Group updated");
  }

  const pubApi = usePublicApi();
  const { data: status } = useAsyncData(async () => {
    const { data } = await pubApi.status();

    return data;
  });

  const { setTheme } = useTheme();

  const auth = useAuthContext();

  const details = computed(() => {
    console.log(auth.user);
    return [
      {
        name: "Name",
        text: auth.user?.name || "Unknown",
      },
      {
        name: "Email",
        text: auth.user?.email || "Unknown",
      },
    ] as Detail[];
  });

  async function deleteProfile() {
    const result = await confirm.open(
      "Are you sure you want to delete your account? If you are the last member in your group all your data will be deleted. This action cannot be undone."
    );

    if (result.isCanceled) {
      return;
    }

    const { response } = await api.user.delete();

    if (response?.status === 204) {
      notify.success("Your account has been deleted.");
      auth.logout(api);
      navigateTo("/");
    }

    notify.error("Failed to delete your account.");
  }

  const token = ref("");
  const tokenUrl = computed(() => {
    if (!window) {
      return "";
    }

    return `${window.location.origin}?token=${token.value}`;
  });

  async function generateToken() {
    const date = new Date();

    const { response, data } = await api.group.createInvitation({
      expiresAt: new Date(date.setDate(date.getDate() + 7)),
      uses: 1,
    });

    if (response?.status === 201) {
      token.value = data.token;
    }
  }

  const passwordChange = reactive({
    loading: false,
    dialog: false,
    current: "",
    new: "",
    isValid: false,
  });

  function openPassChange() {
    passwordChange.dialog = true;
  }

  async function changePassword() {
    passwordChange.loading = true;
    if (!passwordChange.isValid) {
      return;
    }

    const { error } = await api.user.changePassword(passwordChange.current, passwordChange.new);

    if (error) {
      notify.error("Failed to change password.");
      passwordChange.loading = false;
      return;
    }

    notify.success("Password changed successfully.");
    passwordChange.dialog = false;
    passwordChange.new = "";
    passwordChange.current = "";
    passwordChange.loading = false;
  }

  // ===========================================================
  // Notifiers

  const notifiers = useAsyncData(async () => {
    const { data } = await api.notifiers.getAll();

    return data;
  });

  const targetID = ref("");
  const notifier = ref<NotifierCreate | null>(null);
  const notifierDialog = ref(false);

  function openNotifierDialog(v: NotifierOut | null) {
    if (v) {
      targetID.value = v.id;
      notifier.value = {
        name: v.name,
        url: "",
        isActive: v.isActive,
      };
    } else {
      notifier.value = {
        name: "",
        url: "",
        isActive: true,
      };
    }

    notifierDialog.value = true;
  }

  async function createNotifier() {
    if (!notifier.value) {
      return;
    }

    if (targetID.value) {
      await editNotifier();
      return;
    }

    const result = await api.notifiers.create({
      name: notifier.value.name,
      url: notifier.value.url || "",
      isActive: notifier.value.isActive,
    });

    if (result.error) {
      notify.error("Failed to create notifier.");
    }

    notifier.value = null;
    notifierDialog.value = false;

    await notifiers.refresh();
  }

  async function editNotifier() {
    if (!notifier.value) {
      return;
    }

    const result = await api.notifiers.update(targetID.value, {
      name: notifier.value.name,
      url: notifier.value.url || "",
      isActive: notifier.value.isActive,
    });

    if (result.error) {
      notify.error("Failed to update notifier.");
    }

    notifier.value = null;
    notifierDialog.value = false;
    targetID.value = "";

    await notifiers.refresh();
  }

  async function deleteNotifier(id: string) {
    const result = await confirm.open("Are you sure you want to delete this notifier?");

    if (result.isCanceled) {
      return;
    }

    const { error } = await api.notifiers.delete(id);

    if (error) {
      notify.error("Failed to delete notifier.");
      return;
    }

    await notifiers.refresh();
  }

  async function testNotifier() {
    if (!notifier.value) {
      return;
    }

    const { error } = await api.notifiers.test(notifier.value.url);

    if (error) {
      notify.error("Failed to test notifier.");
      return;
    }

    notify.success("Notifier test successful.");
  }
</script>

<template>
  <div>
    <BaseModal v-model="passwordChange.dialog">
      <template #title> Change Password </template>

      <form @submit.prevent="changePassword">
        <FormPassword v-model="passwordChange.current" label="Current Password" placeholder="" />
        <FormPassword v-model="passwordChange.new" label="New Password" placeholder="" />
        <PasswordScore v-model:valid="passwordChange.isValid" :password="passwordChange.new" />

        <div class="flex">
          <BaseButton
            class="ml-auto"
            :loading="passwordChange.loading"
            :disabled="!passwordChange.isValid"
            type="submit"
          >
            Submit
          </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseModal v-model="notifierDialog">
      <template #title> {{ notifier ? "Edit" : "Create" }} Notifier </template>

      <form @submit.prevent="createNotifier">
        <template v-if="notifier">
          <FormTextField v-model="notifier.name" label="Name" />
          <FormTextField v-model="notifier.url" label="URL" />
          <div class="max-w-[100px]">
            <FormCheckbox v-model="notifier.isActive" label="Enabled" />
          </div>
        </template>
        <div class="flex gap-2 justify-between mt-4">
          <BaseButton :disabled="!(notifier && notifier.url)" type="button" @click="testNotifier"> Test </BaseButton>
          <BaseButton type="submit"> Submit </BaseButton>
        </div>
      </form>
    </BaseModal>

    <BaseContainer class="flex flex-col gap-4 mb-6">
      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-account" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> User Profile </span>
            <template #description> Invite users, and manage your account. </template>
          </BaseSectionHeader>
        </template>

        <DetailsSection :details="details" />

        <div class="p-4">
          <div class="flex gap-2">
            <BaseButton size="sm" @click="openPassChange"> Change Password </BaseButton>
            <BaseButton size="sm" @click="generateToken"> Generate Invite Link </BaseButton>
          </div>
          <div v-if="token" class="pt-4 flex items-center pl-1">
            <CopyText class="mr-2 btn-primary btn btn-outline btn-square btn-sm" :text="tokenUrl" />
            {{ tokenUrl }}
          </div>
          <div v-if="token" class="pt-4 flex items-center pl-1">
            <CopyText class="mr-2 btn-primary btn btn-outline btn-square btn-sm" :text="token" />
            {{ token }}
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-megaphone" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Notifiers </span>
            <template #description> Get notifications for up coming maintenance reminders </template>
          </BaseSectionHeader>
        </template>

        <div v-if="notifiers.data.value" class="mx-4 divide-y divide-gray-400 rounded-md border border-gray-400">
          <article v-for="n in notifiers.data.value" :key="n.id" class="p-2">
            <div class="flex flex-wrap items-center gap-2">
              <p class="mr-auto text-lg">{{ n.name }}</p>
              <div class="flex gap-2 justify-end">
                <div class="tooltip" data-tip="Delete">
                  <button class="btn btn-sm btn-square" @click="deleteNotifier(n.id)">
                    <Icon name="mdi-delete" />
                  </button>
                </div>
                <div class="tooltip" data-tip="Edit">
                  <button class="btn btn-sm btn-square" @click="openNotifierDialog(n)">
                    <Icon name="mdi-pencil" />
                  </button>
                </div>
              </div>
            </div>
            <div class="flex justify-between py-1 flex-wrap text-sm">
              <p>
                <span v-if="n.isActive" class="badge badge-success"> Active </span>
                <span v-else class="badge badge-error"> Inactive</span>
              </p>
              <p>
                Created
                <DateTime format="relative" datetime-type="time" :date="n.createdAt" />
              </p>
            </div>
          </article>
        </div>

        <div class="p-4">
          <BaseButton size="sm" @click="openNotifierDialog"> Create </BaseButton>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader class="pb-0">
            <Icon name="mdi-accounts" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Group Settings </span>
            <template #description>
              Shared Group Settings. You may need to refresh your browser for some settings to apply.
            </template>
          </BaseSectionHeader>
        </template>

        <div v-if="group" class="p-5 pt-0">
          <FormSelect v-model="currency" label="Currency Format" :items="currencies" />
          <p class="m-2 text-sm">Example: {{ currencyExample }}</p>

          <div class="mt-4">
            <BaseButton size="sm" @click="updateGroup"> Update Group </BaseButton>
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-fill" class="mr-2 text-base-600" />
            <span class="text-base-600"> Theme Settings </span>
            <template #description>
              Theme settings are stored in your browser's local storage. You can change the theme at any time. If you're
              having trouble setting your theme try refreshing your browser.
            </template>
          </BaseSectionHeader>
        </template>

        <div class="px-4 pb-4">
          <div class="rounded-box grid grid-cols-1 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
            <div
              v-for="theme in themes"
              :key="theme.value"
              class="border-base-content/20 hover:border-base-content/40 outline-base-content overflow-hidden rounded-lg border outline-2 outline-offset-2"
              :data-theme="theme.value"
              :data-set-theme="theme.value"
              data-act-class="outline"
              @click="setTheme(theme.value)"
            >
              <div :data-theme="theme.value" class="bg-base-100 text-base-content w-full cursor-pointer font-sans">
                <div class="grid grid-cols-5 grid-rows-3">
                  <div class="bg-base-200 col-start-1 row-span-2 row-start-1"></div>
                  <div class="bg-base-300 col-start-1 row-start-3"></div>
                  <div class="bg-base-100 col-span-4 col-start-2 row-span-3 row-start-1 flex flex-col gap-1 p-2">
                    <div class="font-bold">{{ theme.label }}</div>
                    <div class="flex flex-wrap gap-1">
                      <div class="bg-primary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-primary-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-secondary flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-secondary-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-accent flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-accent-content text-sm font-bold">A</div>
                      </div>
                      <div class="bg-neutral flex aspect-square w-5 items-center justify-center rounded lg:w-6">
                        <div class="text-neutral-content text-sm font-bold">A</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </BaseCard>

      <BaseCard>
        <template #title>
          <BaseSectionHeader>
            <Icon name="mdi-delete" class="mr-2 -mt-1 text-base-600" />
            <span class="text-base-600"> Delete Account</span>
            <template #description> Delete your account and all its associated data. </template>
          </BaseSectionHeader>
        </template>
        <div class="p-4 px-6 border-t-2 border-gray-300">
          <BaseButton size="sm" class="btn-error" @click="deleteProfile"> Delete Account </BaseButton>
        </div>
      </BaseCard>
    </BaseContainer>
    <footer v-if="status" class="text-center w-full bottom-0 pb-4">
      <p class="text-center text-sm">Version: {{ status.build.version }} ~ Build: {{ status.build.commit }}</p>
    </footer>
  </div>
</template>

<style scoped></style>
